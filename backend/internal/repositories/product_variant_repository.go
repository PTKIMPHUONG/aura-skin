package repositories

import (
	configs "auraskin/internal/configs/firebase"
	"auraskin/internal/databases"
	"auraskin/internal/models"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type ProductVariantRepository interface {
	GetAllVariants() ([]models.ProductVariant, error)
	GetSuggestVariantsForUser(userID string) (*[]models.ProductVariant, error)
	GetSuggestVariantsForAVariant(id string) (*[]models.ProductVariant, error)
	GetVariantByID(id string) (models.ProductVariant, error)
	GetVariantByName(name string) (models.ProductVariant, error)
	CreateVariant(variant models.ProductVariant, productID string) error
	UpdateVariant(id string, variant models.ProductVariant) error
	DeleteVariant(id string) error
	UploadThumbnailAndSetURL(variantID string, file multipart.File, fileHeader *multipart.FileHeader) (string, error)
}

type productVariantRepository struct {
	db *databases.Neo4jDB
}

func NewProductVariantRepository(db *databases.Neo4jDB) ProductVariantRepository {
	return &productVariantRepository{db: db}
}

func (repo *productVariantRepository) GetAllVariants() ([]models.ProductVariant, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (v:ProductVariant) RETURN v", nil)
	if err != nil {
		return nil, err
	}

	var variants []models.ProductVariant
	for result.Next(ctx) {
		record := result.Record()
		node, _ := record.Get("v")
		variantNode := node.(neo4j.Node)

		variantMap := variantNode.Props
		variant, err := (&models.ProductVariant{}).FromMap(variantMap)
		if err != nil {
			return nil, err
		}
		variants = append(variants, *variant)
	}

	return variants, nil
}

func (repo *productVariantRepository) GetSuggestVariantsForAVariant(id string) (*[]models.ProductVariant, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := `
        MATCH (inputPV:ProductVariant {variant_id: $variantID})-[:BELONGS_TO]->(p:Product)-[:BELONGS_TO]->(c:Category)
        MATCH (similarPV:ProductVariant)-[:BELONGS_TO]->(similarP:Product)-[:BELONGS_TO]->(c)
        WHERE similarPV <> inputPV
        RETURN DISTINCT similarPV AS recommendedProductVariant
        LIMIT 4
    `

	params := map[string]interface{}{
		"variantID": id,
	}

	result, err := session.Run(ctx, query, params)
	if err != nil {
		return nil, err
	}

	var variants []models.ProductVariant
	for result.Next(ctx) {
		record := result.Record()
		node, _ := record.Get("recommendedProductVariant")
		variantNode := node.(neo4j.Node)

		variantMap := variantNode.Props
		variant, err := (&models.ProductVariant{}).FromMap(variantMap)
		if err != nil {
			return nil, err
		}
		variants = append(variants, *variant)
	}

	if len(variants) == 0 {
		return nil, errors.New("no recommended products found")
	}

	return &variants, nil
}

func (repo *productVariantRepository) GetSuggestVariantsForUser(userID string) (*[]models.ProductVariant, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := `
        // Dựa trên lịch sử mua hàng
		MATCH (user:User {id: $userID})-[:PLACED_ORDER]->(:Order)-[:CONTAINS]->(purchasedVariant:ProductVariant)-[:BELONGS_TO]->(:Product)-[:BELONGS_TO]->(category:Category)
		WITH category, collect(purchasedVariant) AS purchasedVariants
		MATCH (otherVariant:ProductVariant)-[:BELONGS_TO]->(:Product)-[:BELONGS_TO]->(category)
		WHERE NOT otherVariant IN purchasedVariants
		RETURN DISTINCT otherVariant AS recommendedVariant
		UNION
		// Dựa trên wishlist
		MATCH (user:User {id: $userID})-[:HAS_WISHLIST]->(:Wishlist)-[:WISHES_FOR_VARIANT]->(wishlistVariant:ProductVariant)-[:BELONGS_TO]->(:Product)-[:BELONGS_TO]->(c:Category)
		WITH c
		MATCH (otherVariant:ProductVariant)-[:BELONGS_TO]->(:Product)-[:BELONGS_TO]->(c)
		RETURN DISTINCT otherVariant AS recommendedVariant
		UNION
		// Dựa trên khuyến mãi
		MATCH (sale:Sale)-[:APPLIES_TO_VARIANT]->(discountedVariant:ProductVariant)
		RETURN DISTINCT discountedVariant AS recommendedVariant
    `

	params := map[string]interface{}{
		"userID": userID,
	}

	result, err := session.Run(ctx, query, params)
	if err != nil {
		return nil, err
	}

	var variants []models.ProductVariant
	for result.Next(ctx) {
		record := result.Record()
		node, ok := record.Get("recommendedVariant")
		if !ok {
			continue
		}
		variantNode := node.(neo4j.Node)

		variantMap := variantNode.Props
		variant, err := (&models.ProductVariant{}).FromMap(variantMap)
		if err != nil {
			return nil, err
		}
		variants = append(variants, *variant)
	}

	if len(variants) == 0 {
		return nil, errors.New("no recommended products found")
	}

	return &variants, nil
}

func (repo *productVariantRepository) GetVariantByID(id string) (models.ProductVariant, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (v:ProductVariant {variant_id: $id}) RETURN v", map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return models.ProductVariant{}, err
	}

	if result.Next(ctx) {
		record := result.Record()
		node, found := record.Get("v")
		if !found {
			return models.ProductVariant{}, errors.New("variant not found")
		}
		variantNode := node.(neo4j.Node)

		variantMap := variantNode.Props
		variant, err := (&models.ProductVariant{}).FromMap(variantMap)
		if err != nil {
			return models.ProductVariant{}, err
		}
		return *variant, nil
	}

	return models.ProductVariant{}, errors.New("variant with id " + id + " not found")
}

func (repo *productVariantRepository) GetVariantByName(name string) (models.ProductVariant, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (v:ProductVariant {variant_name: $name}) RETURN v", map[string]interface{}{
		"name": name,
	})
	if err != nil {
		return models.ProductVariant{}, err
	}

	if result.Next(ctx) {
		record := result.Record()
		node, found := record.Get("v")
		if !found {
			return models.ProductVariant{}, errors.New("variant not found")
		}
		variantNode := node.(neo4j.Node)

		variantMap := variantNode.Props
		variant, err := (&models.ProductVariant{}).FromMap(variantMap)
		if err != nil {
			return models.ProductVariant{}, err
		}
		return *variant, nil
	}

	return models.ProductVariant{}, errors.New("variant with name " + name + " not found")
}

func (repo *productVariantRepository) CreateVariant(variant models.ProductVariant, productID string) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Close(ctx)

	// Kiểm tra sự tồn tại của sản phẩm và is_active = true
	productExistsResult, err := tx.Run(ctx,
		"MATCH (p:Product {product_id: $productID, is_active: true}) RETURN p",
		map[string]interface{}{
			"productID": productID,
		},
	)
	if err != nil {
		return err
	}

	if !productExistsResult.Next(ctx) {
		return errors.New("product does not exist or is not active")
	}

	// Tạo node ProductVariant
	variantMap := variant.ToMap()
	variantMap["is_active"] = true

	_, err = tx.Run(ctx,
		"CREATE (v:ProductVariant {variant_id: $variant_id, variant_name: $variant_name, size: $size, color: $color, price: $price, stock_quantity: $stock_quantity, thumbnail: $thumbnail, is_active: $is_active, description_images: $description_images})",
		variantMap,
	)
	if err != nil {
		return err
	}

	// Tạo relationship BELONGS_TO giữa ProductVariant và Product
	_, err = tx.Run(ctx,
		`
        MATCH (p:Product {product_id: $productID}), (v:ProductVariant {variant_id: $variant_id})
        MERGE (v)-[r:BELONGS_TO]->(p)
        RETURN r
        `,
		map[string]interface{}{
			"productID":  productID,
			"variant_id": variantMap["variant_id"],
		},
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (repo *productVariantRepository) UpdateVariant(id string, variant models.ProductVariant) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Close(ctx)

	variantMap := variant.ToMap()

	// Kiểm tra variant có tồn tại hay không
	variantExistsResult, err := tx.Run(ctx,
		"MATCH (v:ProductVariant {variant_id: $variant_id}) RETURN v",
		map[string]interface{}{
			"variant_id": id,
		},
	)
	if err != nil {
		return err
	}

	if !variantExistsResult.Next(ctx) {
		return errors.New("variant does not exist")
	}

	_, err = tx.Run(ctx,
		"MATCH (v:ProductVariant {variant_id: $variant_id}) SET v.variant_name = $variant_name, v.size = $size, v.color = $color, v.price = $price, v.stock_quantity = $stock_quantity, v.thumbnail = $thumbnail, v.is_active = $is_active, v.description_images: $description_images RETURN v",
		variantMap,
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (repo *productVariantRepository) DeleteVariant(id string) error {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Close(ctx)

	_, err = tx.Run(ctx,
		"MATCH (v:ProductVariant {variant_id: $variant_id, is_active: true}) SET v.is_active = false RETURN v",
		map[string]interface{}{
			"variant_id": id,
		},
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
func (repo *productVariantRepository) UploadThumbnailAndSetURL(variantID string, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	ctx := context.Background()
	session := repo.db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), fileHeader.Filename)

	client, err := configs.InitializeFirebaseApp().Storage(ctx)
	if err != nil {
		return "", err
	}

	bucket, err := client.Bucket(configs.FirebaseStorageBucketName)
	if err != nil {
		return "", err
	}
	writer := bucket.Object(fileName).NewWriter(ctx)
	token := generateUUID()
	writer.Metadata = map[string]string{
		"firebaseStorageDownloadTokens": token,
	}
	defer writer.Close()

	if _, err := io.Copy(writer, file); err != nil {
		return "", err
	}
	encodedFileName := url.QueryEscape(fileName)

	fileUrl := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s", configs.FirebaseStorageBucketName, encodedFileName, token)

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Close(ctx)

	_, err = tx.Run(ctx,
		"MATCH (v:ProductVariant {variant_id: $variant_id}) SET v.thumbnail = $thumbnail_url RETURN v",
		map[string]interface{}{
			"variant_id":    variantID,
			"thumbnail_url": fileUrl,
		},
	)
	if err != nil {
		return "", err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", err
	}

	return fileUrl, nil
}