// Copyright 2020 Decipher Technology Studios
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package render

//import (
//	. "github.com/smartystreets/goconvey/convey"
//)

// func TestParseVariable(t *testing.T) {

// 	Convey("When render.ParseVariable is invoked", t, func() {

// 		name := tests.MustGenerateHex(t)
// 		value := tests.MustGenerateHex(t)

// 		Convey("with an empty name", func() {

// 			name := ""
// 			variable, err := ParseVariable(fmt.Sprintf("%s=%s", name, value))

// 			Convey("it should return a variable", func() {

// 				Convey("with and empty name", func() {
// 					So(variable.Name, ShouldBeZeroValue)
// 				})

// 				Convey("with the expected value", func() {
// 					So(variable.Value, ShouldEqual, value)
// 				})
// 			})

// 			Convey("it should return a nil error", func() {
// 				So(err, ShouldBeNil)
// 			})
// 		})

// 		Convey("with an empty value", func() {

// 			value := ""
// 			variable, err := ParseVariable(fmt.Sprintf("%s=%s", name, value))

// 			Convey("it should return a variable", func() {

// 				Convey("with the expected name", func() {
// 					So(variable.Name, ShouldEqual, name)
// 				})

// 				Convey("with an empty value", func() {
// 					So(variable.Value, ShouldBeZeroValue)
// 				})
// 			})

// 			Convey("it should return a nil error", func() {
// 				So(err, ShouldBeNil)
// 			})
// 		})

// 		Convey("with too few equal signs", func() {

// 			variable, err := ParseVariable(fmt.Sprintf("%s%s", name, value))

// 			Convey("it should return a nil variable", func() {
// 				So(variable, ShouldBeNil)
// 			})

// 			Convey("it should return a non-nil error", func() {
// 				So(err, ShouldNotBeNil)
// 			})
// 		})

// 		Convey("with too many equal signs", func() {

// 			variable, err := ParseVariable(fmt.Sprintf("=%s=%s=", name, value))

// 			Convey("it should return a nil variable", func() {
// 				So(variable, ShouldBeNil)
// 			})

// 			Convey("it should return a non-nil error", func() {
// 				So(err, ShouldNotBeNil)
// 			})
// 		})
// 	})
// }

// func TestParseVariables(t *testing.T) {

// 	Convey("When render.ParseVariables is invoked", t, func() {

// 		variable := &Variable{Name: tests.MustGenerateHex(t), Value: tests.MustGenerateHex(t)}

// 		Convey("with an invalid variables", func() {

// 			variables, err := ParseVariables([]string{fmt.Sprintf("%s%s", variable.Name, variable.Value)})

// 			Convey("it should return a nil slice", func() {
// 				So(variables, ShouldBeNil)
// 			})

// 			Convey("it should return a non-nil error", func() {
// 				So(err, ShouldNotBeNil)
// 			})
// 		})

// 		Convey("with valid variables", func() {

// 			variables, err := ParseVariables([]string{fmt.Sprintf("%s=%s", variable.Name, variable.Value)})

// 			Convey("it should return the expected slice", func() {
// 				So(variables, ShouldEqual, []*Variable{variable})
// 			})

// 			Convey("it should return a nil error", func() {
// 				So(err, ShouldBeNil)
// 			})
// 		})
// 	})
// }
