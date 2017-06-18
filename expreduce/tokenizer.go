// CAUTION: Generated file - DO NOT EDIT.

package expreduce

import (
	"fmt"
	"log"
	"math/big"
	"strings"
)

type Calclexer struct {
	s       string
	pos     int
	buf     []rune
	empty   bool
	current rune
}

func newLexer(s string) (y *Calclexer) {
	y = &Calclexer{s: s}
	if y.pos != len(y.s) {
		y.current = rune(y.s[y.pos])
	}
	/*fmt.Printf("y.current: %d, y.pos: %d, '%s'\n", y.current, y.pos, y.buf)*/
	y.pos += 1
	return
}

func (y *Calclexer) getc() rune {
	if y.current != 0 {
		y.buf = append(y.buf, y.current)
	}
	y.current = 0
	if y.pos != len(y.s) {
		y.current = rune(y.s[y.pos])
	}
	/*fmt.Printf("y.current: %d, y.pos: %d, '%s'\n", y.current, y.pos, y.buf)*/
	y.pos += 1
	return y.current
}

func (y Calclexer) Error(e string) {
	fmt.Printf("Syntax::sntx: %v.\n\n\n", e)
}

func (y *Calclexer) Lex(lval *CalcSymType) int {
	/*var err error*/
	c := y.current
	if y.empty {
		c, y.empty = y.getc(), false
	}

	/*E  [eE][-+]?{D}*/
	/*F  {D}"."{D}?{E}?|{D}{E}?|"."{D}{E}?*/

yystate0:

	y.buf = y.buf[:0]

	goto yystart1

	goto yystate0 // silence unused label error
	goto yystate1 // silence unused label error
yystate1:
	c = y.getc()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate3
	case c == '"':
		goto yystate5
	case c == '#':
		goto yystate9
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	case c == '&':
		goto yystate18
	case c == '(':
		goto yystate20
	case c == ')':
		goto yystate21
	case c == '*':
		goto yystate22
	case c == '+':
		goto yystate23
	case c == ',':
		goto yystate24
	case c == '-':
		goto yystate25
	case c == '.':
		goto yystate27
	case c == '/':
		goto yystate31
	case c == ':':
		goto yystate38
	case c == ';':
		goto yystate42
	case c == '<':
		goto yystate44
	case c == '=':
		goto yystate47
	case c == '>':
		goto yystate50
	case c == '?':
		goto yystate52
	case c == '@':
		goto yystate53
	case c == '[':
		goto yystate55
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate2
	case c == ']':
		goto yystate56
	case c == '^':
		goto yystate57
	case c == '_':
		goto yystate11
	case c == '{':
		goto yystate58
	case c == '|':
		goto yystate59
	case c == '}':
		goto yystate61
	case c >= '0' && c <= '9':
		goto yystate37
	}

yystate2:
	c = y.getc()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate2
	}

yystate3:
	c = y.getc()
	switch {
	default:
		goto yyrule33
	case c == '=':
		goto yystate4
	}

yystate4:
	c = y.getc()
	goto yyrule35

yystate5:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate6
	case c == '\\':
		goto yystate7
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate5
	}

yystate6:
	c = y.getc()
	goto yyrule4

yystate7:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate8
	case c == '\\':
		goto yystate7
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate5
	}

yystate8:
	c = y.getc()
	switch {
	default:
		goto yyrule4
	case c == '"':
		goto yystate6
	case c == '\\':
		goto yystate7
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate5
	}

yystate9:
	c = y.getc()
	goto yyrule36

yystate10:
	c = y.getc()
	switch {
	default:
		goto yyrule51
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	case c == '_':
		goto yystate11
	case c == '`':
		goto yystate17
	}

yystate11:
	c = y.getc()
	switch {
	default:
		goto yyrule50
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate12
	case c == '.':
		goto yystate14
	case c == '_':
		goto yystate15
	}

yystate12:
	c = y.getc()
	switch {
	default:
		goto yyrule50
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate12
	case c == '`':
		goto yystate13
	}

yystate13:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate14:
	c = y.getc()
	goto yyrule50

yystate15:
	c = y.getc()
	switch {
	default:
		goto yyrule50
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate12
	case c == '.':
		goto yystate14
	case c == '_':
		goto yystate16
	}

yystate16:
	c = y.getc()
	switch {
	default:
		goto yyrule50
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate12
	case c == '.':
		goto yystate14
	}

yystate17:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate18:
	c = y.getc()
	switch {
	default:
		goto yyrule34
	case c == '&':
		goto yystate19
	}

yystate19:
	c = y.getc()
	goto yyrule48

yystate20:
	c = y.getc()
	goto yyrule5

yystate21:
	c = y.getc()
	goto yyrule6

yystate22:
	c = y.getc()
	goto yyrule18

yystate23:
	c = y.getc()
	goto yyrule16

yystate24:
	c = y.getc()
	goto yyrule7

yystate25:
	c = y.getc()
	switch {
	default:
		goto yyrule17
	case c == '>':
		goto yystate26
	}

yystate26:
	c = y.getc()
	goto yyrule14

yystate27:
	c = y.getc()
	switch {
	default:
		goto yyrule30
	case c == '.':
		goto yystate28
	case c >= '0' && c <= '9':
		goto yystate30
	}

yystate28:
	c = y.getc()
	switch {
	default:
		goto yyrule31
	case c == '.':
		goto yystate29
	}

yystate29:
	c = y.getc()
	goto yyrule32

yystate30:
	c = y.getc()
	switch {
	default:
		goto yyrule3
	case c >= '0' && c <= '9':
		goto yystate30
	}

yystate31:
	c = y.getc()
	switch {
	default:
		goto yyrule19
	case c == '.':
		goto yystate32
	case c == '/':
		goto yystate33
	case c == ';':
		goto yystate35
	case c == '@':
		goto yystate36
	}

yystate32:
	c = y.getc()
	goto yyrule37

yystate33:
	c = y.getc()
	switch {
	default:
		goto yyrule39
	case c == '.':
		goto yystate34
	}

yystate34:
	c = y.getc()
	goto yyrule38

yystate35:
	c = y.getc()
	goto yyrule24

yystate36:
	c = y.getc()
	goto yyrule43

yystate37:
	c = y.getc()
	switch {
	default:
		goto yyrule2
	case c == '.':
		goto yystate30
	case c >= '0' && c <= '9':
		goto yystate37
	}

yystate38:
	c = y.getc()
	switch {
	default:
		goto yyrule46
	case c == ':':
		goto yystate39
	case c == '=':
		goto yystate40
	case c == '>':
		goto yystate41
	}

yystate39:
	c = y.getc()
	goto yyrule49

yystate40:
	c = y.getc()
	goto yyrule22

yystate41:
	c = y.getc()
	goto yyrule15

yystate42:
	c = y.getc()
	switch {
	default:
		goto yyrule8
	case c == ';':
		goto yystate43
	}

yystate43:
	c = y.getc()
	goto yyrule9

yystate44:
	c = y.getc()
	switch {
	default:
		goto yyrule27
	case c == '=':
		goto yystate45
	case c == '>':
		goto yystate46
	}

yystate45:
	c = y.getc()
	goto yyrule29

yystate46:
	c = y.getc()
	goto yyrule42

yystate47:
	c = y.getc()
	switch {
	default:
		goto yyrule21
	case c == '=':
		goto yystate48
	}

yystate48:
	c = y.getc()
	switch {
	default:
		goto yyrule25
	case c == '=':
		goto yystate49
	}

yystate49:
	c = y.getc()
	goto yyrule23

yystate50:
	c = y.getc()
	switch {
	default:
		goto yyrule26
	case c == '=':
		goto yystate51
	}

yystate51:
	c = y.getc()
	goto yyrule28

yystate52:
	c = y.getc()
	goto yyrule44

yystate53:
	c = y.getc()
	switch {
	default:
		goto yyrule40
	case c == '@':
		goto yystate54
	}

yystate54:
	c = y.getc()
	goto yyrule41

yystate55:
	c = y.getc()
	goto yyrule10

yystate56:
	c = y.getc()
	goto yyrule11

yystate57:
	c = y.getc()
	goto yyrule20

yystate58:
	c = y.getc()
	goto yyrule12

yystate59:
	c = y.getc()
	switch {
	default:
		goto yyrule45
	case c == '|':
		goto yystate60
	}

yystate60:
	c = y.getc()
	goto yyrule47

yystate61:
	c = y.getc()
	goto yyrule13

yyrule1: // [ \t\r]+

	goto yystate0
yyrule2: // {D}
	{

		var base int = 10
		tmpi := big.NewInt(0)
		_, ok := tmpi.SetString(string(y.buf), base)
		if !ok {
			log.Fatal("Failed in integer parsing.")
		}
		lval.val = &Integer{tmpi}
		return INTEGER
	}
yyrule3: // {F}
	{

		tmpf := big.NewFloat(0)
		_, ok := tmpf.SetString(string(y.buf))
		if !ok {
			log.Fatal("Failed in float parsing.")
		}
		lval.val = &Flt{tmpf}
		return FLOAT
	}
yyrule4: // {S}
	{

		tmps := string(y.buf)
		lval.val = &String{tmps[1 : len(tmps)-1]}
		return STRING
	}
yyrule5: // \(
	{
		return LPARSYM /* skipped */
	}
yyrule6: // \)
	{
		return RPARSYM /* skipped */
	}
yyrule7: // ,
	{
		return COMMASYM /* skipped */
	}
yyrule8: // ;
	{
		return SEMISYM /* CompoundExpression */
	}
yyrule9: // ;;
	{
		return SPANSYM /* Span */
	}
yyrule10: // \[
	{
		return LBRACKETSYM /* skipped */
	}
yyrule11: // \]
	{
		return RBRACKETSYM /* skipped */
	}
yyrule12: // \{
	{
		return LCURLYSYM /* skipped */
	}
yyrule13: // \}
	{
		return RCURLYSYM /* skipped */
	}
yyrule14: // ->
	{
		return RULESYM /* Rule */
	}
yyrule15: // :>
	{
		return RULEDELAYEDSYM /* RuleDelayed */
	}
yyrule16: // \+
	{
		return PLUSSYM /* Plus */
	}
yyrule17: // \-
	{
		return MINUSSYM /* */
	}
yyrule18: // \*
	{
		return MULTSYM /* */
	}
yyrule19: // \/
	{
		return DIVSYM /* */
	}
yyrule20: // \^
	{
		return EXPSYM /* */
	}
yyrule21: // =
	{
		return SETSYM /* set */
	}
yyrule22: // :=
	{
		return SETDELAYEDSYM /* setdelayed */
	}
yyrule23: // ===
	{
		return SAMESYM /* SameQ */
	}
yyrule24: // \/;
	{
		return CONDITIONSYM /* Condition */
	}
yyrule25: // ==
	{
		return EQUALSYM /* Equal */
	}
yyrule26: // >
	{
		return GREATERSYM /* Greater */
	}
yyrule27: // \<
	{
		return LESSSYM /* Less */
	}
yyrule28: // >=
	{
		return GREATEREQUALSYM /* GreaterEqual */
	}
yyrule29: // \<=
	{
		return LESSEQUALSYM /* LessEqual */
	}
yyrule30: // \.
	{
		return DOTSYM /* Dot, Optional */
	}
yyrule31: // \.\.
	{
		return REPEATEDSYM /* Repeated */
	}
yyrule32: // \.\.\.
	{
		return REPEATEDNULLSYM /* RepeatedNull */
	}
yyrule33: // !
	{
		return EXCLAMATIONSYM /* Factorial, Not */
	}
yyrule34: // &
	{
		return FUNCTIONSYM /* Factorial */
	}
yyrule35: // !=
	{
		return UNEQUALSYM /* Unequal */
	}
yyrule36: // #
	{
		return SLOTSYM /* Slot */
	}
yyrule37: // \/\.
	{
		return REPLACEALLSYM /* ReplaceAll */
	}
yyrule38: // \/\/\.
	{
		return REPLACEREPSYM /* ReplaceRepeated */
	}
yyrule39: // \/\/
	{
		return POSTFIXSYM /* */
	}
yyrule40: // @
	{
		return FUNCAPPSYM /* */
	}
yyrule41: // @@
	{
		return APPLYSYM /* */
	}
yyrule42: // \<>
	{
		return STRINGJOINSYM /* StringJoin */
	}
yyrule43: // \/@
	{
		return MAPSYM /* Map */
	}
yyrule44: // \?
	{
		return PATTESTSYM /* PatternTest */
	}
yyrule45: // \|
	{
		return ALTSYM /* Alternatives */
	}
yyrule46: // :
	{
		return COLONSYM /* Pattern */
	}
yyrule47: // \|\|
	{
		return ORSYM /* Or */
	}
yyrule48: // &&
	{
		return ANDSYM /* And */
	}
yyrule49: // ::
	{
		return MESSAGENAMESYM /* MessageName */
	}
yyrule50: // {pattern}
	{

		delim := "_"
		blankType := &Symbol{"Blank"}
		if strings.Contains(string(y.buf), "___") {
			delim = "___"
			blankType = &Symbol{"BlankNullSequence"}
		} else if strings.Contains(string(y.buf), "__") {
			delim = "__"
			blankType = &Symbol{"BlankSequence"}
		}
		parts := strings.Split(string(y.buf), delim)
		if len(parts) == 1 {
			lval.val = NewExpression([]Ex{&Symbol{"Pattern"}, &Symbol{parts[0]}, NewExpression([]Ex{blankType})})
			return PATTERN
		}
		if len(parts) == 2 {
			if parts[0] == "" {
				if parts[1] == "" {
					lval.val = NewExpression([]Ex{blankType})
				} else if delim == "_" && parts[1] == "." {
					lval.val = NewExpression([]Ex{&Symbol{"Optional"}, NewExpression([]Ex{blankType})})
				} else {
					lval.val = NewExpression([]Ex{blankType, &Symbol{parts[1]}})
				}
				return PATTERN
			} else {
				if parts[1] == "" {
					lval.val = NewExpression([]Ex{&Symbol{"Pattern"}, &Symbol{parts[0]}, NewExpression([]Ex{blankType})})
				} else if delim == "_" && parts[1] == "." {
					lval.val = NewExpression([]Ex{&Symbol{"Optional"}, NewExpression([]Ex{&Symbol{"Pattern"}, &Symbol{parts[0]}, NewExpression([]Ex{blankType})})})
				} else {
					lval.val = NewExpression([]Ex{&Symbol{"Pattern"}, &Symbol{parts[0]}, NewExpression([]Ex{blankType, &Symbol{parts[1]}})})
				}
				return PATTERN
			}
		}
		lval.val = NewExpression([]Ex{&Symbol{"Error"}, &String{"Pattern parse error."}})
		return PATTERN
	}
yyrule51: // {contextedIdent}
	{

		lval.val = &Symbol{string(y.buf)}
		return NAME
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	y.empty = true
	return int(c)
}
