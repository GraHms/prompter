package prompter

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPromptBuilder(t *testing.T) {
	Convey("Given a fresh PromptBuilder", t, func() {
		builder := New()

		Convey("When no messages have been added", func() {
			result := builder.Build()
			Convey("Then Build should return an empty string", func() {
				So(result, ShouldEqual, "")
			})
		})

		Convey("When a system message is added", func() {
			builder := New().System("initialize system")
			result := builder.Build()
			Convey("Then the output should start with the system block", func() {
				So(result, ShouldStartWith, "<|system|>\ninitialize system")
			})
		})

		Convey("When user and assistant messages are added", func() {
			builder := New().
				User("hello world").
				Assistant("")
			result := builder.Build()
			Convey("Then the prompt should contain both blocks in order", func() {
				want := "<|user|>\nhello world\n\n<|assistant|>\n\n\n"
				So(result, ShouldEqual, want)
			})
		})

		Convey("When a custom role is added via Role()", func() {
			const RoleValidator Role = "validator"
			builder := New().Role(RoleValidator, "check this")
			result := builder.Build()
			Convey("Then the custom block should appear correctly", func() {
				want := "<|validator|>\ncheck this\n\n"
				So(result, ShouldEqual, want)
			})
		})

		Convey("When multiple messages with different roles are added", func() {
			builder := New().
				System("sys").
				User("usr").
				Role("loader", "load data").
				Assistant("resp")
			result := builder.Build()

			want := "<|system|>\nsys\n\n" +
				"<|user|>\nusr\n\n" +
				"<|loader|>\nload data\n\n" +
				"<|assistant|>\nresp\n\n"
			Convey("Then Build should preserve order and formatting exactly", func() {
				So(result, ShouldEqual, want)
			})
		})

		Convey("When user content contains multiple lines", func() {
			multi := "line1\nline2\nline3"
			builder := New().User(multi)
			result := builder.Build()
			Convey("Then all lines should appear unaltered in the prompt", func() {
				So(result, ShouldContainSubstring, "line1\nline2\nline3")
			})
		})
	})
}
