package program

import (
	"database/sql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gym-app/common/logger"
	"regexp"
)

var _ = Describe("Repository", func() {
	var repository *PRepository
	var mock sqlmock.Sqlmock

	BeforeEach(func() {
		var db *sql.DB
		var err error

		// db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) // use equal matcher
		db, mock, err = sqlmock.New() // mock sql.DB
		Expect(err).ShouldNot(HaveOccurred())

		gdb, err := gorm.Open(postgres.New(postgres.Config{
			Conn: db,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
		Expect(err).ShouldNot(HaveOccurred())

		var prRepository = NewPRepository(gdb, logger.NewLogger())
		repository = &prRepository
	})
	AfterEach(func() {
		err := mock.ExpectationsWereMet() // make sure all expectations were met
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("list all", func() {
		It("empty", func() {
			const sqlSelectAll = `SELECT * FROM "programs"`
			mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAll)).
				WillReturnRows(sqlmock.NewRows(nil))

			l := repository.Get()
			Expect(l).Should(BeEmpty())
		})
	})

	//Context("search", func() {
	//	It("from date", func() {
	//		rows := sqlmock.
	//			NewRows([]string{"id", "title", "content", "tags", "created_at"}).
	//			AddRow(1, "post 1", "hello 1", nil, time.Now())
	//
	//		// limit/offset is not parameter
	//		const sqlSearch = `
	//			SELECT * FROM "blogs"
	//			WHERE (title like $1)
	//			LIMIT 10 OFFSET 0`
	//		const q = "os"
	//
	//		mock.ExpectQuery(regexp.QuoteMeta(sqlSearch)).
	//			WithArgs("%" + q + "%").
	//			WillReturnRows(rows)
	//
	//		l, err := repository.SearchByTitle(q, 0, 10)
	//		Expect(err).ShouldNot(HaveOccurred())
	//
	//		Expect(l).Should(HaveLen(1))
	//		Expect(l[0].Title).Should(ContainSubstring(q))
	//	})
	//})
})
