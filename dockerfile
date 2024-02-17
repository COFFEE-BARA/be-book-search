# 사용할 Go 버전 지정
FROM golang:1.20 AS builder

# 작업 디렉토리 설정
WORKDIR /app

# 의존성 파일 복사 및 다운로드
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# 소스 코드 복사
COPY . .

# 애플리케이션 빌드
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 최종 이미지
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/main .
CMD ["./main"]
