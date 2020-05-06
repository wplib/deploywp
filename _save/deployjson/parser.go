package deployjson

type parser struct {
	text string
	len  int
	pos  int
}

func newParser(s string) *parser {
	return &parser{
		text: s,
		len:  len(s),
	}
}

func (me *parser) canParse() bool {
	return me.pos < me.len
}

func (me *parser) eat(ch byte) (found bool) {
	for me.pos < me.len {
		if byte(me.text[me.pos]) != ch {
			me.pos++
			continue
		}
		found = true
		me.pos++
		break
	}
	return found
}

func (me *parser) count(ch byte) (c int) {
	c = 0
	for byte(me.text[me.pos]) == ch {
		me.pos++
		if me.pos == me.len {
			c = -1
		}
		c++
	}
	return c
}

func (me *parser) captureTo(ch byte) string {
	b := make([]byte, 0)
	for me.pos < me.len {
		if byte(me.text[me.pos]) == ch {
			break
		}
		b = append(b, me.text[me.pos])
		me.pos++
	}
	return string(b)
}
