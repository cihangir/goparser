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
		file, err = ParseFile("file_test.go")
	})

	Context("When parsing successfull", func() {
		It("should return false", func() {
			Expect(false).To(BeFalse())
		})
		It("should not error", func() {
			Expect(err).NotTo(HaveOccured())
		})
		XIt("should have file name", func() {
			Expect(file.Name).To(Equal("file_test.go"))
		})
		XIt("should consist path", func() {
			Expect(file.Name).To(Equal("file_test.go"))
		})
		It("should consist package name", func() {
			Expect(file.PackageName).To(Equal("goparser_test"))
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
		Context("when functions have parameters", func() {
			Context("when functions have incoming parameters", func() {
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
				It("function should have outgoing parameters type property", func() {
					for _, fun := range file.Functions {
						if fun.Name == "MethodWith2Input2Output" {
							Expect(fun.OutgoingParams[0].TypeOf).To(Equal("string"))
						}
					}
				})

			})
			Context("when functions have outgoing parameters", func() {
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
			})
		})

		Context("when functions doesnt have parameters", func() {
			It("should not have inccmoing parameters", func() {
				for _, fun := range file.Functions {
					if fun.Name == "MethodWithNoInputNoOutput" {
						Expect(len(fun.IncomingParams)).To(Equal(0))
					}
				}
			})
			It("should not have outgoing parameters", func() {
				for _, fun := range file.Functions {
					if fun.Name == "MethodWithNoInputNoOutput" {
						Expect(len(fun.OutgoingParams)).To(Equal(0))
					}
				}
			})
		})

		It("for functions name property should be empty string", func() {
			for _, fun := range file.Functions {
				if fun.Name == "MethodWith2Input2Output" {
					Expect(fun.OutgoingParams[0].Name).To(Equal(""))
				}
			}
		})
		Context("when file have structs", func() {
			It("should have more than one struct", func() {
				structs := file.Structs
				Expect(len(structs)).NotTo(BeZero())
			})
			It("should have name", func() {
				structs := file.Structs[0]
				Expect(structs.Name).To(Equal("Filter"))
			})
			It("should have fields", func() {
				structs := file.Structs[0]
				Expect(len(structs.Fields)).NotTo(Equal(0))
			})
			Context("when struct has fields", func() {
				It("should have more than one field", func() {
					structs := file.Structs[0]
					Expect(len(structs.Fields)).NotTo(Equal(0))
				})
				It("should have more than one field", func() {
					structs := file.Structs[0]
					Expect(len(structs.Fields)).NotTo(Equal(0))
				})
				var stt *StructField
				BeforeEach(func() {
					stt = file.Structs[0].Fields[0]
				})
				Context("when struct has field", func() {
					It("should have name", func() {
						Expect(stt.Name).To(Equal("minAge"))
					})
					It("should have type", func() {
						Expect(stt.TypeOf).To(Equal("int"))
					})
					It("should have documentation", func() {
						Expect(stt.Documentation).To(ContainSubstring("@Data"))
					})
				})
			})
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
