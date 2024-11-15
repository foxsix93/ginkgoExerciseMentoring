package bank_test

import (
	"exercise_mentoringGO/bank"
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Bank Account", func() {

	var account bank.Account

	ginkgo.BeforeEach(func() {
		account = bank.Account{}
	})

	ginkgo.Describe("Deposit", func() {
		ginkgo.It("should deposit money and increase the balance", func() {
			err := account.Deposit(100)
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(account.GetBalance()).To(gomega.Equal(100.0))
		})

		ginkgo.It("should return an error if depositing zero or negative amount", func() {
			err := account.Deposit(0)
			gomega.Expect(err).To(gomega.HaveOccurred())
			gomega.Expect(account.GetBalance()).To(gomega.Equal(0.0))

			err = account.Deposit(-50)
			gomega.Expect(err).To(gomega.HaveOccurred())
			gomega.Expect(account.GetBalance()).To(gomega.Equal(0.0))
		})
	})

	ginkgo.Describe("Withdraw", func() {
		ginkgo.It("should withdraw money and decrease the balance", func() {
			account.Deposit(200)
			err := account.Withdraw(50)
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(account.GetBalance()).To(gomega.Equal(150.0))
		})

		ginkgo.It("should return an error if withdrawing more than the balance", func() {
			account.Deposit(100)
			err := account.Withdraw(150)
			gomega.Expect(err).To(gomega.HaveOccurred())
			gomega.Expect(account.GetBalance()).To(gomega.Equal(100.0))
		})

		ginkgo.It("should return an error if withdrawing zero or negative amount", func() {
			account.Deposit(100)

			err := account.Withdraw(0)
			gomega.Expect(err).To(gomega.HaveOccurred())
			gomega.Expect(account.GetBalance()).To(gomega.Equal(100.0))

			err = account.Withdraw(-50)
			gomega.Expect(err).To(gomega.HaveOccurred())
			gomega.Expect(account.GetBalance()).To(gomega.Equal(100.0))
		})
	})

	ginkgo.Describe("GetBalance", func() {
		ginkgo.It("should return the correct balance", func() {
			account.Deposit(100)
			account.Withdraw(50)
			gomega.Expect(account.GetBalance()).To(gomega.Equal(50.0))
		})
	})

})

func TestAccount(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Bank Account Suite")
}
