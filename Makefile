install_dependencies:
	@echo "Installing dependencies..."
	@go get -u \
    github.com/gin-gonic/gin \
    github.com/golang-jwt/jwt/v4 \
    github.com/joho/godotenv \
    golang.org/x/crypto \
    gorm.io/driver/postgres \
    gorm.io/gorm

	@echo "Done"