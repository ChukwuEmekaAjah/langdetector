package langdetector

import (
	"reflect"
	"testing"
)

func TestNewScanner(t *testing.T) {
	type args struct {
		spec LangSpec
	}
	tests := []struct {
		name string
		args args
		want *Scanner
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScanner(tt.args.spec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScanner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_WithKeywords(t *testing.T) {
	type fields struct {
		keywords          []string
		identifierPattern string
		source            string
		line              uint
		current           uint
		errors            []error
		language          string
	}
	type args struct {
		keywords []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Scanner
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				keywords:          tt.fields.keywords,
				identifierPattern: tt.fields.identifierPattern,
				source:            tt.fields.source,
				line:              tt.fields.line,
				current:           tt.fields.current,
				errors:            tt.fields.errors,
				language:          tt.fields.language,
			}
			if got := s.WithKeywords(tt.args.keywords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scanner.WithKeywords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_WithSource(t *testing.T) {
	type fields struct {
		keywords          []string
		identifierPattern string
		source            string
		line              uint
		current           uint
		errors            []error
		language          string
	}
	type args struct {
		source string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Scanner
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				keywords:          tt.fields.keywords,
				identifierPattern: tt.fields.identifierPattern,
				source:            tt.fields.source,
				line:              tt.fields.line,
				current:           tt.fields.current,
				errors:            tt.fields.errors,
				language:          tt.fields.language,
			}
			if got := s.WithSource(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scanner.WithSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_WithSourcefile(t *testing.T) {
	type fields struct {
		keywords          []string
		identifierPattern string
		source            string
		line              uint
		current           uint
		errors            []error
		language          string
	}
	type args struct {
		sourceFile string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Scanner
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				keywords:          tt.fields.keywords,
				identifierPattern: tt.fields.identifierPattern,
				source:            tt.fields.source,
				line:              tt.fields.line,
				current:           tt.fields.current,
				errors:            tt.fields.errors,
				language:          tt.fields.language,
			}
			got, err := s.WithSourcefile(tt.args.sourceFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scanner.WithSourcefile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scanner.WithSourcefile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_ScanTokens(t *testing.T) {
	type fields struct {
		keywords          []string
		identifierPattern string
		source            string
		line              uint
		current           uint
		errors            []error
		language          string
	}
	tests := []struct {
		name   string
		fields fields
		want   []Token
	}{{
		name: "Scan 1",
		fields: fields{
			source:            "+=++%===23.34ajah abstract",
			language:          "js",
			identifierPattern: "^[a-zA-Z_$]+[0-9a-zA-Z_$]*",
			keywords: []string{
				"abstract", "arguments", "await", "boolean",
				"break", "byte", "case", "catch",
				"char", "class", "const", "continue",
				"debugger", "default", "delete", "do",
				"double", "else", "enum", "eval",
				"export", "extends", "false", "final",
				"finally", "float", "for", "function",
				"goto", "if", "implements", "import",
				"in", "instanceof", "int", "interface",
				"let", "long", "native", "new",
				"null", "package", "private", "protected",
				"public", "return", "short", "static",
				"super", "switch", "synchronized", "this",
				"throw", "throws", "transient", "true",
				"try", "typeof", "var", "void",
				"volatile", "while", "with", "yield",
			},
		},
		want: []Token{{
			tokenType: "PLUS_ASSIGN",
			literal:   "+=",
			lexeme:    "+=",
		}, {
			tokenType: "PLUS_PLUS",
			literal:   "++",
			lexeme:    "++",
		}, {
			tokenType: "MODULO_ASSIGN",
			literal:   "%=",
			lexeme:    "%=",
		}, {
			tokenType: "EQUAL_EQUAL",
			literal:   "==",
			lexeme:    "==",
		}, {
			tokenType: "NUMBER",
			literal:   "23.34",
			lexeme:    "23.34",
		}, {
			tokenType: "IDENTIFIER",
			literal:   "ajah",
			lexeme:    "ajah",
		}, {
			tokenType: "KEYWORD",
			literal:   "abstract",
			lexeme:    "abstract",
		}},
	}, {
		name: "Scan 2",
		fields: fields{
			source:            "var x = 22; class Ajah{}",
			language:          "js",
			identifierPattern: "^[a-zA-Z_$]+[0-9a-zA-Z_$]*",
			keywords: []string{
				"abstract", "arguments", "await", "boolean",
				"break", "byte", "case", "catch",
				"char", "class", "const", "continue",
				"debugger", "default", "delete", "do",
				"double", "else", "enum", "eval",
				"export", "extends", "false", "final",
				"finally", "float", "for", "function",
				"goto", "if", "implements", "import",
				"in", "instanceof", "int", "interface",
				"let", "long", "native", "new",
				"null", "package", "private", "protected",
				"public", "return", "short", "static",
				"super", "switch", "synchronized", "this",
				"throw", "throws", "transient", "true",
				"try", "typeof", "var", "void",
				"volatile", "while", "with", "yield",
			},
		},
		want: []Token{{
			tokenType: "KEYWORD",
			literal:   "var",
			lexeme:    "var",
		}, {
			tokenType: "IDENTIFIER",
			literal:   "x",
			lexeme:    "x",
		}, {
			tokenType: "ASSIGN",
			literal:   "=",
			lexeme:    "=",
		}, {
			tokenType: "NUMBER",
			literal:   "22",
			lexeme:    "22",
		}, {
			tokenType: "SEMI_COLON",
			literal:   ";",
			lexeme:    ";",
		}, {
			tokenType: "KEYWORD",
			literal:   "class",
			lexeme:    "class",
		}, {
			tokenType: "IDENTIFIER",
			literal:   "Ajah",
			lexeme:    "Ajah",
		}, {
			tokenType: "LEFT_BRACE",
			literal:   "{",
			lexeme:    "{",
		}, {
			tokenType: "RIGHT_BRACE",
			literal:   "}",
			lexeme:    "}",
		}},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				keywords:          tt.fields.keywords,
				identifierPattern: tt.fields.identifierPattern,
				source:            tt.fields.source,
				line:              tt.fields.line,
				current:           tt.fields.current,
				errors:            tt.fields.errors,
				language:          tt.fields.language,
			}
			// println(s.ScanTokens())
			if got := s.ScanTokens(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scanner.ScanTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_advance(t *testing.T) {
	type fields struct {
		keywords          []string
		identifierPattern string
		source            string
		line              uint
		current           uint
		errors            []error
		language          string
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				keywords:          tt.fields.keywords,
				identifierPattern: tt.fields.identifierPattern,
				source:            tt.fields.source,
				line:              tt.fields.line,
				current:           tt.fields.current,
				errors:            tt.fields.errors,
				language:          tt.fields.language,
			}
			if got := s.advance(); got != tt.want {
				t.Errorf("Scanner.advance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_isAtEnd(t *testing.T) {
	type fields struct {
		keywords          []string
		identifierPattern string
		source            string
		line              uint
		current           uint
		errors            []error
		language          string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Scanner{
				keywords:          tt.fields.keywords,
				identifierPattern: tt.fields.identifierPattern,
				source:            tt.fields.source,
				line:              tt.fields.line,
				current:           tt.fields.current,
				errors:            tt.fields.errors,
				language:          tt.fields.language,
			}
			if got := s.isAtEnd(); got != tt.want {
				t.Errorf("Scanner.isAtEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_addToken(t *testing.T) {
	type fields struct {
		keywords          []string
		identifierPattern string
		source            string
		line              uint
		current           uint
		start             uint
		errors            []error
		language          string
		tokens            []Token
		score             int
	}
	type args struct {
		tokenType string
		lexeme    string
		literal   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				keywords:          tt.fields.keywords,
				identifierPattern: tt.fields.identifierPattern,
				source:            tt.fields.source,
				line:              tt.fields.line,
				current:           tt.fields.current,
				start:             tt.fields.start,
				errors:            tt.fields.errors,
				language:          tt.fields.language,
				tokens:            tt.fields.tokens,
				score:             tt.fields.score,
			}
			s.addToken(tt.args.tokenType, tt.args.lexeme, tt.args.literal)
		})
	}
}

func TestScanner_isSpecialChar(t *testing.T) {
	type fields struct {
		keywords          []string
		identifierPattern string
		source            string
		line              uint
		current           uint
		start             uint
		errors            []error
		language          string
		tokens            []Token
		score             int
	}
	type args struct {
		ch byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Scanner{
				keywords:          tt.fields.keywords,
				identifierPattern: tt.fields.identifierPattern,
				source:            tt.fields.source,
				line:              tt.fields.line,
				current:           tt.fields.current,
				start:             tt.fields.start,
				errors:            tt.fields.errors,
				language:          tt.fields.language,
				tokens:            tt.fields.tokens,
				score:             tt.fields.score,
			}
			if got := s.isSpecialChar(tt.args.ch); got != tt.want {
				t.Errorf("Scanner.isSpecialChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_isAlpha(t *testing.T) {
	type fields struct {
		keywords          []string
		identifierPattern string
		source            string
		line              uint
		current           uint
		start             uint
		errors            []error
		language          string
		tokens            []Token
		score             int
	}
	type args struct {
		ch byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Scanner{
				keywords:          tt.fields.keywords,
				identifierPattern: tt.fields.identifierPattern,
				source:            tt.fields.source,
				line:              tt.fields.line,
				current:           tt.fields.current,
				start:             tt.fields.start,
				errors:            tt.fields.errors,
				language:          tt.fields.language,
				tokens:            tt.fields.tokens,
				score:             tt.fields.score,
			}
			if got := s.isAlpha(tt.args.ch); got != tt.want {
				t.Errorf("Scanner.isAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_identifier(t *testing.T) {
	type fields struct {
		keywords          []string
		identifierPattern string
		source            string
		line              uint
		current           uint
		start             uint
		errors            []error
		language          string
		tokens            []Token
		score             int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				keywords:          tt.fields.keywords,
				identifierPattern: tt.fields.identifierPattern,
				source:            tt.fields.source,
				line:              tt.fields.line,
				current:           tt.fields.current,
				start:             tt.fields.start,
				errors:            tt.fields.errors,
				language:          tt.fields.language,
				tokens:            tt.fields.tokens,
				score:             tt.fields.score,
			}
			s.identifier()
		})
	}
}

func TestScanner_isNumber(t *testing.T) {
	type fields struct {
		keywords          []string
		identifierPattern string
		source            string
		line              uint
		current           uint
		start             uint
		errors            []error
		language          string
		tokens            []Token
		score             int
	}
	type args struct {
		ch byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Scanner{
				keywords:          tt.fields.keywords,
				identifierPattern: tt.fields.identifierPattern,
				source:            tt.fields.source,
				line:              tt.fields.line,
				current:           tt.fields.current,
				start:             tt.fields.start,
				errors:            tt.fields.errors,
				language:          tt.fields.language,
				tokens:            tt.fields.tokens,
				score:             tt.fields.score,
			}
			if got := s.isNumber(tt.args.ch); got != tt.want {
				t.Errorf("Scanner.isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_number(t *testing.T) {
	type fields struct {
		keywords          []string
		identifierPattern string
		source            string
		line              uint
		current           uint
		start             uint
		errors            []error
		language          string
		tokens            []Token
		score             int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				keywords:          tt.fields.keywords,
				identifierPattern: tt.fields.identifierPattern,
				source:            tt.fields.source,
				line:              tt.fields.line,
				current:           tt.fields.current,
				start:             tt.fields.start,
				errors:            tt.fields.errors,
				language:          tt.fields.language,
				tokens:            tt.fields.tokens,
				score:             tt.fields.score,
			}
			s.number()
		})
	}
}

func TestScanner_string(t *testing.T) {
	type fields struct {
		keywords          []string
		identifierPattern string
		source            string
		line              uint
		current           uint
		start             uint
		errors            []error
		language          string
		tokens            []Token
		score             int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				keywords:          tt.fields.keywords,
				identifierPattern: tt.fields.identifierPattern,
				source:            tt.fields.source,
				line:              tt.fields.line,
				current:           tt.fields.current,
				start:             tt.fields.start,
				errors:            tt.fields.errors,
				language:          tt.fields.language,
				tokens:            tt.fields.tokens,
				score:             tt.fields.score,
			}
			s.string()
		})
	}
}

func TestScanner_match(t *testing.T) {
	type fields struct {
		keywords          []string
		identifierPattern string
		source            string
		line              uint
		current           uint
		start             uint
		errors            []error
		language          string
		tokens            []Token
		score             int
	}
	type args struct {
		ch byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				keywords:          tt.fields.keywords,
				identifierPattern: tt.fields.identifierPattern,
				source:            tt.fields.source,
				line:              tt.fields.line,
				current:           tt.fields.current,
				start:             tt.fields.start,
				errors:            tt.fields.errors,
				language:          tt.fields.language,
				tokens:            tt.fields.tokens,
				score:             tt.fields.score,
			}
			if got := s.match(tt.args.ch); got != tt.want {
				t.Errorf("Scanner.match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_peek(t *testing.T) {
	type fields struct {
		keywords          []string
		identifierPattern string
		source            string
		line              uint
		current           uint
		start             uint
		errors            []error
		language          string
		tokens            []Token
		score             int
	}
	type args struct {
		position uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Scanner{
				keywords:          tt.fields.keywords,
				identifierPattern: tt.fields.identifierPattern,
				source:            tt.fields.source,
				line:              tt.fields.line,
				current:           tt.fields.current,
				start:             tt.fields.start,
				errors:            tt.fields.errors,
				language:          tt.fields.language,
				tokens:            tt.fields.tokens,
				score:             tt.fields.score,
			}
			if got := s.peek(tt.args.position); got != tt.want {
				t.Errorf("Scanner.peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_GetScoring(t *testing.T) {
	langsSpec := []LangSpec{
		{
			language: "js",
			keywords: []string{
				"abstract", "arguments", "await", "boolean",
				"break", "byte", "case", "catch",
				"char", "class", "const", "continue",
				"debugger", "default", "delete", "do",
				"double", "else", "enum", "eval",
				"export", "extends", "false", "final",
				"finally", "float", "for", "function",
				"goto", "if", "implements", "import",
				"in", "instanceof", "int", "interface",
				"let", "long", "native", "new",
				"null", "package", "private", "protected",
				"public", "return", "short", "static",
				"super", "switch", "synchronized", "this",
				"throw", "throws", "transient", "true",
				"try", "typeof", "var", "void", "require",
				"volatile", "while", "with", "yield",
			},
			identifierPattern: "^[a-zA-Z_$]+[0-9a-zA-Z_$]*",
		},
		{
			language: "java",
			keywords: []string{
				"abstract", "assert", "boolean", "break", "byte",
				"case", "catch", "char", "class", "const", "continue",
				"default", "double", "do", "else", "enum", "extends",
				"false", "final", "finally", "float", "for", "goto",
				"if", "implements", "import", "instanceof", "int",
				"interface", "long", "native", "new", "null", "package",
				"private", "protected", "public", "return", "short",
				"static", "strictfp", "super", "switch", "synchronized",
				"this", "throw", "throws", "transient", "true", "try",
				"void", "volatile", "while",
			},
			identifierPattern: "^[a-zA-Z_$]+[0-9a-zA-Z_$]*",
		},
		{
			language: "golang",
			keywords: []string{
				"const", "chan", "break", "defer", "var", "interface", "case",
				"go", "func", "map", "continue", "type", "struct", "default",
				"import", "else", "package", "fallthrough", "for", "goto", "if",
				"range", "return", "select", "switch",
			},
			identifierPattern: "^[a-zA-Z_$]+[0-9a-zA-Z_$]*",
		},
	}
	code := `
	const mongoose = require("mongoose");
	const bcrypt = require("bcrypt");
	
	const authSchema = new mongoose.Schema({
		firstName: { type: String, default: "" },
		lastName: { type: String, default: "" },
		password: { type: String },
		email: { type: String, unique: true },
		verified: { type: Boolean, default: false },
		lastSignin: { type: Date, default: Date.now },
		organizationId: { type: String, default: "" },
		role: { type: String, default: "admin" },
		meta: { type: Object, default: {} },
	});
	
	authSchema.virtual("name").get(function () {
		return ${this.firstName} ${this.lastName};
	});
	
	authSchema.statics.generateHash = function (password) {
		return bcrypt.hashSync(password, bcrypt.genSaltSync(10));
	};
	
	authSchema.methods.compareHash = function (password) {
		return bcrypt.compareSync(password, this.password);
	};
	
	module.exports = {
		name: "Researcher",
		schema: authSchema,
	};
	`
	for _, spec := range langsSpec {
		scanner := NewScanner(spec)
		scanner.WithSource(code)
		scanner.ScanTokens()
		t.Logf("score for language: %s is: %d \n", spec.language, scanner.GetScore())
	}
}
