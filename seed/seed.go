package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/badzboss/go-elasticsearch/models"
)

// ข้อมูลสินค้าสำหรับสุ่ม - เหมาะกับการค้นหาด้วย Elasticsearch
var products = []struct {
	name        string
	category    string
	description string
}{
	{
		name:        "MacBook Pro 16-inch M3 Max",
		category:    "Electronics",
		description: "Apple MacBook Pro with M3 Max chip, 16-inch Liquid Retina XDR display, 36GB RAM, 1TB SSD. Perfect for professional developers and content creators. Features outstanding performance, long battery life, and brilliant display quality.",
	},
	{
		name:        "Sony WH-1000XM5 Wireless Headphones",
		category:    "Electronics",
		description: "Industry-leading noise canceling wireless headphones with premium sound quality. 30-hour battery life, multipoint connection, and exceptional comfort. Ideal for music lovers and professionals who need to focus.",
	},
	{
		name:        "Nike Air Max 270",
		category:    "Fashion",
		description: "Nike's iconic sneakers with Max Air cushioning unit. Provides exceptional all-day comfort with a modern, stylish design. Available in multiple colors. Perfect for casual wear and light exercise.",
	},
	{
		name:        "Samsung Galaxy S24 Ultra",
		category:    "Electronics",
		description: "Samsung's flagship smartphone with 200MP camera, S Pen support, and powerful Snapdragon processor. Features 6.8-inch Dynamic AMOLED display, 12GB RAM, and 5G connectivity. Perfect for photography and productivity.",
	},
	{
		name:        "IKEA MALM Bed Frame",
		category:    "Furniture",
		description: "Modern bed frame with clean lines and timeless design. Made from sustainable wood materials. Available in queen and king sizes. Easy to assemble with storage options available.",
	},
	{
		name:        "Instant Pot Duo 7-in-1",
		category:    "Home & Kitchen",
		description: "Multi-functional electric pressure cooker that can pressure cook, slow cook, rice cooker, steamer, sauté, yogurt maker, and food warmer. Saves time and energy with smart programming.",
	},
	{
		name:        "Canon EOS R5 Camera",
		category:    "Electronics",
		description: "Professional full-frame mirrorless camera with 45MP sensor and 8K video recording. Features advanced autofocus, in-body image stabilization, and weather sealing. Perfect for professional photographers and videographers.",
	},
	{
		name:        "The North Face Thermoball Jacket",
		category:    "Fashion",
		description: "Insulated jacket with synthetic ThermoBall insulation that mimics down. Lightweight, compressible, and maintains warmth even when wet. Ideal for outdoor adventures and cold weather.",
	},
	{
		name:        "Dyson V15 Detect Vacuum Cleaner",
		category:    "Home & Kitchen",
		description: "Cordless vacuum cleaner with laser dust detection and intelligent suction. HEPA filtration captures 99.99% of particles. Powerful cleaning performance for all floor types.",
	},
	{
		name:        "Herman Miller Aeron Chair",
		category:    "Furniture",
		description: "Ergonomic office chair with PostureFit SL support and 8Z Pellicle suspension. Designed for all-day comfort and proper posture. Industry-leading 12-year warranty. Perfect for home office and professional workspace.",
	},
	{
		name:        "Kindle Paperwhite",
		category:    "Electronics",
		description: "E-reader with 6.8-inch glare-free display and adjustable warm light. Waterproof design with weeks of battery life. Access to millions of books. Perfect for avid readers.",
	},
	{
		name:        "Levi's 501 Original Fit Jeans",
		category:    "Fashion",
		description: "Classic straight-leg jeans with iconic button fly. Made from premium denim with a timeless fit. Available in various washes. The original blue jean since 1873.",
	},
	{
		name:        "Vitamix Professional Blender",
		category:    "Home & Kitchen",
		description: "Professional-grade blender with powerful motor and aircraft-grade stainless steel blades. Makes smoothies, hot soups, frozen desserts, and more. Self-cleaning feature and 7-year warranty.",
	},
	{
		name:        "Nintendo Switch OLED",
		category:    "Electronics",
		description: "Gaming console with vibrant 7-inch OLED screen and enhanced audio. Plays in TV mode, tabletop mode, or handheld mode. Access to thousands of games including Mario, Zelda, and Pokemon.",
	},
	{
		name:        "West Elm Mid-Century Sofa",
		category:    "Furniture",
		description: "Modern sofa with clean lines and tapered legs inspired by 1950s design. Available in multiple fabrics and colors. Kiln-dried hardwood frame ensures lasting durability.",
	},
	{
		name:        "Patagonia Better Sweater Jacket",
		category:    "Fashion",
		description: "Quarter-zip fleece jacket made from recycled polyester. Provides warmth without weight. Durable water-repellent finish. Fair Trade Certified sewn. Perfect layer for outdoor activities.",
	},
	{
		name:        "KitchenAid Stand Mixer",
		category:    "Home & Kitchen",
		description: "Professional 5-quart stand mixer with 10-speed control. Includes whisk, dough hook, and flat beater. Powerful enough for bread dough yet gentle enough for whipped cream. Available in 40+ colors.",
	},
	{
		name:        "Apple Watch Series 9",
		category:    "Electronics",
		description: "Smartwatch with advanced health features including heart rate monitoring, ECG, and blood oxygen measurement. Features always-on Retina display and up to 18-hour battery life. Perfect fitness companion.",
	},
	{
		name:        "Zara Linen Blend Blazer",
		category:    "Fashion",
		description: "Lightweight blazer made from linen blend fabric. Perfect for summer business casual or smart casual occasions. Features notched lapels and front pockets. Available in multiple colors.",
	},
	{
		name:        "Le Creuset Cast Iron Dutch Oven",
		category:    "Home & Kitchen",
		description: "Enameled cast iron Dutch oven perfect for slow cooking, braising, and baking. Superior heat distribution and retention. Oven-safe up to 500°F. Lifetime warranty. Available in signature colors.",
	},
}

func main() {
	// เชื่อมต่อกับ database
	models.ConnectDatabase()

	// ตั้งค่า random seed
	rand.Seed(time.Now().UnixNano())

	log.Println("Starting to seed blog data...")

	// สร้างข้อมูล 10,000 บล็อก
	batchSize := 100
	totalBlogs := 10000

	for i := 0; i < totalBlogs; i++ {
		// สุ่มสินค้า
		product := products[rand.Intn(len(products))]

		// เพิ่ม variation เพื่อให้มีความหลากหลาย
		title := fmt.Sprintf("%s - %s Edition #%d", product.name, product.category, rand.Intn(1000)+1)

		// สร้าง body ที่มี keyword หลากหลายเพื่อการค้นหา
		body := fmt.Sprintf("Category: %s\n\n%s\n\n", product.category, product.description)
		body += generateAdditionalDetails(product.category)

		blog := models.Blog{
			Title: title,
			Body:  body,
		}

		// Insert ข้อมูล
		if err := models.DB.Create(&blog).Error; err != nil {
			log.Printf("Error creating blog %d: %v", i+1, err)
			continue
		}

		// แสดง progress ทุก 100 records
		if (i+1)%batchSize == 0 {
			log.Printf("Created %d/%d blogs...", i+1, totalBlogs)
		}
	}

	log.Printf("✅ Successfully seeded %d blogs!", totalBlogs)
}

// สร้างรายละเอียดเพิ่มเติมตาม category
func generateAdditionalDetails(category string) string {
	details := ""

	switch category {
	case "Electronics":
		features := []string{
			"High-performance processor with advanced technology",
			"Extended battery life for all-day use",
			"Premium build quality with attention to detail",
			"Latest connectivity options including WiFi 6 and Bluetooth 5.0",
			"Enhanced security features for data protection",
		}
		details += "Key Features:\n"
		// สุ่ม 2-3 features
		numFeatures := rand.Intn(2) + 2
		for i := 0; i < numFeatures; i++ {
			details += fmt.Sprintf("• %s\n", features[rand.Intn(len(features))])
		}

	case "Fashion":
		styles := []string{
			"Modern and trendy design suitable for any occasion",
			"Premium quality materials for comfort and durability",
			"Available in multiple sizes and color options",
			"Easy care instructions for long-lasting wear",
			"Versatile style that pairs well with various outfits",
		}
		details += "Style & Care:\n"
		numStyles := rand.Intn(2) + 2
		for i := 0; i < numStyles; i++ {
			details += fmt.Sprintf("• %s\n", styles[rand.Intn(len(styles))])
		}

	case "Furniture":
		specs := []string{
			"Sturdy construction with quality materials",
			"Easy assembly with included instructions",
			"Fits perfectly in modern and traditional spaces",
			"Available in multiple finishes and sizes",
			"Backed by manufacturer warranty",
		}
		details += "Specifications:\n"
		numSpecs := rand.Intn(2) + 2
		for i := 0; i < numSpecs; i++ {
			details += fmt.Sprintf("• %s\n", specs[rand.Intn(len(specs))])
		}

	case "Home & Kitchen":
		benefits := []string{
			"Makes meal preparation faster and easier",
			"Durable construction for years of reliable use",
			"Easy to clean and maintain",
			"Energy efficient design saves money",
			"Versatile functionality for various cooking needs",
		}
		details += "Benefits:\n"
		numBenefits := rand.Intn(2) + 2
		for i := 0; i < numBenefits; i++ {
			details += fmt.Sprintf("• %s\n", benefits[rand.Intn(len(benefits))])
		}
	}

	// เพิ่มราคาสุ่ม
	price := rand.Intn(9000) + 1000
	details += fmt.Sprintf("\nPrice: $%d.00\n", price)
	details += fmt.Sprintf("Stock: %d units available\n", rand.Intn(100)+1)

	// เพิ่ม tags สำหรับการค้นหา
	details += "\nTags: "
	tags := []string{"popular", "best-seller", "new-arrival", "featured", "trending", "limited-edition"}
	numTags := rand.Intn(3) + 2
	for i := 0; i < numTags; i++ {
		if i > 0 {
			details += ", "
		}
		details += tags[rand.Intn(len(tags))]
	}
	details += "\n"

	return details
}
