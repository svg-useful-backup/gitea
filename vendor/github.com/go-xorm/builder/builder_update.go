package builder

import (
	"errors"
	"fmt"
)

func (b *Builder) updateWriteTo(w Writer) error {
	if len(b.tableName) <= 0 {
		return errors.New("no table indicated")
	}
	if len(b.updates) <= 0 {
		return errors.New("no column to be update")
	}

	if _, err := fmt.Fprintf(w, "UPDATE %s SET ", b.tableName); err != nil {
		return err
	}

	for i, s := range b.updates {
		if err := s.opWriteTo(",", w); err != nil {
			return err
		}

		if i != len(b.updates)-1 {
			if _, err := fmt.Fprint(w, ","); err != nil {
				return err
			}
		}
	}

	if _, err := fmt.Fprint(w, " WHERE "); err != nil {
		return err
	}

	return b.cond.WriteTo(w)
}
