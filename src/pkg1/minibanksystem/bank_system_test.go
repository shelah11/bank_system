package minibanksystem_test

import (
	"banksystem/src/pkg1/minibanksystem"
	"encoding/json"
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MiniBankSys", func() {
	var a minibanksystem.Testing
	BeforeEach(func() {
		a = minibanksystem.Testing{
			Name:          " ",
			CurrentAmount: 0,
			Deposit:       0,
			Withdraws:     0,
		}
	})
	Describe("GetCurrentAccount", func() {
		Context("when GetCurrentAmount function is called", func() {
			It("should return the current amount in the customer account", func() {
				a = minibanksystem.Testing{
					CurrentAmount: 0,
				}
				Expect(a.GetCurrentAmount()).To(Equal(float32(0)))
			})
		})
	})

	Describe("AddIntoAccount", func() {
		Context("when AddIntoAccount function is called", func() {
			It("should return the current amount in the customer account", func() {
				a.CurrentAmount = 10
				deposit := 50
				text := "deposited"
				amount, msg := a.AddIntoAccount(float32(deposit))

				Expect(amount).To(Equal(float32(deposit)))
				Expect(msg).To(Equal(text))
				Expect(a.CurrentAmount).To(Equal(float32(60)))
			})
		})
	})

	Describe("WithdrawsMoney", func() {
		Context("when WithdrawsMoney function is called", func() {
			It("should return the amount you withdrew", func() {
				a.CurrentAmount = 50
				withdraw := 50.0
				withdraw_amount, err := a.WithdrawsMoney(float32(withdraw))

				Expect(withdraw_amount).To(Equal(float32(0)))
				Expect(err).ShouldNot(HaveOccurred())
				Expect(a.CurrentAmount).To(Equal(float32(0)))
			})

			It("should return an err if you do not have enough money", func() {
				a.CurrentAmount = 10
				withdraw := 60.0
				_, err := a.WithdrawsMoney(float32(withdraw))

				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(errors.New("You do not have enough")))
				Expect(a.CurrentAmount).To(Equal(float32(10)))
			})
		})
	})

	Describe("OpenNewAccount", func() {
		Context("when OpenNewAccount function is called", func() {
			It("should create an new account with customer name and deposit ", func() {
				deposit := float32(50.0)
				name := "test"
				customer, _ := a.OpenNewAccount(name, float32(deposit))

				var reading minibanksystem.Testing
				err := json.Unmarshal([]byte(customer), &reading)

				Expect(err).ToNot(HaveOccurred())
				Expect(reading.CurrentAmount).To(Equal(deposit))
				Expect(reading.Deposit).To(Equal(deposit))
				Expect(reading.Name).To(Equal(name))
			})
			It("should return an error if name is empty", func() {
				deposit := 0.0
				name := ""
				_, err := a.OpenNewAccount(name, float32(deposit))

				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(errors.New("make sure you provide a valid name")))
			})
			It("should return an error if name is is numeric or alphanumeric", func() {
				deposit := 0.0
				name := "llk1234"
				_, err := a.OpenNewAccount(name, float32(deposit))

				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(errors.New("name must only contain letters")))
			})
			It("should return an error if deposit is zero", func() {
				deposit := 0.0
				name := "test"
				_, err := a.OpenNewAccount(name, float32(deposit))

				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(errors.New("a minimum of £20 is required")))
			})
			It("should return an error if deposit is greater than zero but less than 20", func() {
				deposit := 12.0
				name := "test"
				_, err := a.OpenNewAccount(name, float32(deposit))

				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(errors.New("a minimum of £20 is required")))
			})

		})
	})
})
