package main

import (
	"errors"
	"fmt"
)

/*
========================================
        FORMAT VERBS (QUICK GUIDE)
========================================

%v   → default format (uses String() if exists)
%+v  → struct with field names (but overridden by String())
%#v  → Go syntax (raw struct, ignores String())
%T   → type
%s   → string
%d   → integer
%f   → float (%.2f for 2 decimal places)
%t   → boolean
%q   → quoted string
%%   → print %

========================================
*/

type ConfigItem struct {
	Key   string
	Value interface{}
	IsSet bool
}

/*
========================================

	STRING METHOD

========================================

- This overrides default printing
- Used by %v and %+v
- Makes output cleaner

NOTE:
Use %v for Value because it can be any type
========================================
*/
func (c ConfigItem) String() string {
	return fmt.Sprintf("Key: %s, Value: %v, IsSet: %t", c.Key, c.Value, c.IsSet)
}

func main() {

	// Basic variables
	appName := "EnvParser"
	version := 1.2
	port := 8080
	isEnabled := true

	/*
		========================================
		    FORMATTED STRING EXAMPLE
		========================================
	*/
	status := fmt.Sprintf(
		"Application: %s (Version: %.1f) running on port %d. Enabled: %t",
		appName, version, port, isEnabled,
	)
	fmt.Println(status)

	// Creating struct values
	item1 := ConfigItem{Key: "API_URL", Value: "http://localhost:3000/api", IsSet: true}
	item2 := ConfigItem{Key: "TIMEOUT_MS", Value: 5000, IsSet: true}
	item3 := ConfigItem{Key: "DEBUG_MODE", Value: false, IsSet: false}

	/*
		========================================
		    PRINTING DIFFERENT FORMATS
		========================================
	*/

	// %v → uses String() method
	fmt.Printf("Item 1 (%%v): %v\n", item1)

	// %+v → also uses String() (so same output)
	fmt.Printf("Item 2 (%%+v): %+v\n", item2)

	// %#v → shows raw Go struct (ignores String())
	fmt.Printf("Item 3 (%%#v): %#v\n", item3)

	/*
		========================================
		        ERROR FORMATTING
		========================================
	*/
	err := errors.New("test")

	// %w → wraps error (used for error chaining)
	fmt.Errorf("here is the error on port %d: %w", port, err)
}


