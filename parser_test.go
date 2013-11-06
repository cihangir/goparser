package goparser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/siesta/goparser"
	"testing"
)

func TestGovalidator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goparser Suite")
}

var _ = Describe("Reading file", func() {
	var (
		file *ParsedFile
		err  error
	)

	BeforeEach(func() {
		file, err = ParseFile("test.go")
	})

	Context("When parsing successfull", func() {
		It("should return false", func() {
			Expect(false).To(BeFalse())
		})
		It("should not error", func() {
			Expect(err).NotTo(HaveOccured())
		})
		XIt("should have file name", func() {
			Expect(file.Name).To(Equal("test.go"))
		})
		XIt("should consist path", func() {
			Expect(file.Name).To(Equal("test.go"))
		})
		It("should consist package name", func() {
			Expect(file.PackageName).To(Equal("goparser"))
		})
		It("should consist 2 imports", func() {
			Expect(len(file.Imports)).To(Equal(2))
		})
		It("should consist fmt imports", func() {
			Expect(file.Imports[0].Path).To(ContainSubstring("fmt"))
		})
		It("should consist time imports", func() {
			Expect(file.Imports[1].Path).To(ContainSubstring("time"))
		})
		It("should have 12 number of functions", func() {
			Expect(len(file.Functions)).To(Equal(13))
		})

		It("function should have name property", func() {
			firstFunction := file.Functions[0]
			Expect(firstFunction.Name).To(ContainSubstring("MethodWithNoInputNoOutput"))
		})
		It("function should have Documentation property", func() {
			firstFunction := file.Functions[0]
			Expect(firstFunction.Documentation).To(ContainSubstring("s is documentation for MethodWithNoInputNoOutput"))
		})
		It("function should have Documentation property", func() {
			firstFunction := file.Functions[0]
			Expect(firstFunction.Receiver).To(ContainSubstring("Filter"))
		})

		It("function should have imcoming parameters property", func() {
			for _, fun := range file.Functions {
				if fun.Name == "MethodWith2Input2Output" {
					Expect(len(fun.IncomingParams)).To(Equal(2))
				}
			}
		})
		It("function should have imcoming firstParam property", func() {
			for _, fun := range file.Functions {
				if fun.Name == "MethodWith2Input2Output" {
					Expect(fun.IncomingParams[0].Name).To(ContainSubstring("firstParam"))
				}
			}
		})

		It("function should have imcoming secondParam property", func() {
			for _, fun := range file.Functions {
				if fun.Name == "MethodWith2Input2Output" {
					Expect(fun.IncomingParams[1].Name).To(ContainSubstring("secondParam"))
				}
			}
		})

		It("function should have outgoing parameters property", func() {
			for _, fun := range file.Functions {
				if fun.Name == "MethodWith2Input2Output" {
					Expect(len(fun.OutgoingParams)).To(Equal(2))
				}
			}
		})

		It("function should have imcoming parameters type property", func() {
			for _, fun := range file.Functions {
				if fun.Name == "MethodWith2Input2Output" {
					Expect(fun.IncomingParams[0].TypeOf).To(Equal("string"))
				}
			}
		})

		It("function should have outgoing parameters type property", func() {
			for _, fun := range file.Functions {
				if fun.Name == "MethodWith2Input2Output" {
					Expect(fun.OutgoingParams[0].TypeOf).To(Equal("string"))
				}
			}
		})

		It("function, name property should be empty string", func() {
			for _, fun := range file.Functions {
				if fun.Name == "MethodWith2Input2Output" {
					Expect(fun.OutgoingParams[0].Name).To(Equal(""))
				}
			}
		})
	})

	Context("when parsing fails", func() {

		BeforeEach(func() {
			//this file is not there
			file, err = ParseFile("test_2.go")
		})

		It("should return the zero-value for the file", func() {
			Expect(file).To(BeZero())
		})

		It("should error", func() {
			Expect(err).To(HaveOccured())
		})
	})

})

// var _ = Describe("Parsing whole content", func() {
// 	var (
// 		file *ParsedFile
// 		err  error
// 	)

// 	BeforeEach(func() {
// 		file, err = ParseFile("test.go")
// 	})

// 	Context("When parsing successfull", func() {
// 		It("should return false", func() {
// 			Expect(false).To(BeFalse())
// 		})
// 		It("should not error", func() {
// 			Expect(err).NotTo(HaveOccured())
// 		})
// 	})

// })
