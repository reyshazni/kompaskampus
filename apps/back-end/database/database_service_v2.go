package database

import (
	"FindMyDosen/model/entity"
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var databasev2 *gorm.DB

func loadDBv2() {
	tlsConf := createTLSConf()
	if err := mysql.RegisterTLSConfig("custom", &tlsConf); err != nil {
		log.Fatal(err)
		return
	}
	dsn := os.Getenv("DSN")
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return
	}
	gormDB, err := gorm.Open(gormmysql.New(gormmysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	gormDB.AutoMigrate(&entity.University{})
	gormDB.AutoMigrate(&entity.Lecture{})
	gormDB.AutoMigrate(&entity.Subject{})
	gormDB.AutoMigrate(&entity.LectureSubject{})
	gormDB.AutoMigrate(&entity.User{})
	gormDB.AutoMigrate(&entity.RefreshToken{})
	gormDB.AutoMigrate(&entity.LectureRating{})
	databasev2 = gormDB
}

func GetDBv2() *gorm.DB {
	once.Do(loadDBv2)
	return databasev2
}

func createTLSConf() tls.Config {
	rootCertPool := x509.NewCertPool()
	pem, err := os.ReadFile("cert/ca-cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append PEM.")
	}
	clientCert := make([]tls.Certificate, 0, 1)

	certs, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	clientCert = append(clientCert, certs)

	return tls.Config{
		RootCAs:            rootCertPool,
		Certificates:       clientCert,
		InsecureSkipVerify: true, // needed for self-signed certs
	}
}
