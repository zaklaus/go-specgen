/*
   Copyright 2019 Dominik Madarász
   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at
       http://www.apache.org/licenses/LICENSE-2.0
   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package gospec

import (
	"bytes"
	"fmt"
	"strings"
)

// ExportGo exports data to the Go language format
func (ctx *Context) ExportGo() {
	fmt.Printf("/* This file has been generated by go-specgen */\npackage %s\n\n", ctx.FormatName)

	for _, en := range ctx.Enums {
		fmt.Printf("/* %s */\nconst (\n", en.Name)
		for _, ev := range en.Fields {
			fmt.Printf("\t%s ", ev.Value)

			if ev.DocString != "" {
				fmt.Printf(" /* %s */", strings.TrimSpace(ev.DocString))
			}

			fmt.Println("")
		}
		fmt.Printf(")\n\n")
	}

	fmt.Println("")

	for _, spec := range ctx.Specs {
		if spec.DocString != "" {
			fmt.Printf("/* %s */\n", strings.TrimSpace(spec.DocString))
		}

		fmt.Printf("type %s struct {\n", spec.Name)
		for _, fld := range spec.Fields {
			dumpField(fld)
		}
		fmt.Printf("}\n\n")
	}
}

func dumpField(field Field) {
	fmt.Printf("\t%s ", field.Name)

	if field.IsPointer {
		fmt.Print("*")
	} else if field.IsArray {
		dumpArrayGo(field, field)
	}

	if field.InnerArray == nil {
		fmt.Printf("%s ", field.Type)
	}

	if len(field.Tags) > 0 {
		var buf bytes.Buffer
		for _, tag := range field.Tags {
			buf.WriteString(fmt.Sprintf("%s:\"%s\" ",
				tag.Name,
				strings.TrimSuffix(strings.Join(tag.Values, ","), ","),
			))
		}

		fmt.Printf(" `%s` ", strings.TrimSpace(buf.String()))
	}

	if field.DocString != "" {
		fmt.Printf(" /* %s */", strings.TrimSpace(field.DocString))
	}

	fmt.Println("")
}

func dumpArrayGo(field, orig Field) {
	if field.ArrayLen == 0 {
		fmt.Print("[]")
	} else if field.ArrayLen > 0 {
		fmt.Printf("[%d]", field.ArrayLen)
	}

	if field.InnerArray != nil {
		dumpArrayGo(*field.InnerArray, field)
	} else if orig.Name != field.Name {
		fmt.Printf("%s ", field.Type)
	}
}
