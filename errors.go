package validator

import "errors"

var (
	errStringToInt          = errors.New("validator: Die Zeichenkette konnte nicht in eine Ganzzahl umgewandelt werden")
	errStringToFloat        = errors.New("validator: Die Zeichenkette konnte nicht in eine Gleitkommazahl umgewandelt werden")
	errRequireRules         = errors.New("validator: Bitte geben Sie mindestens eine Regel für die Validate*-Methode an")
	errValidateArgsMismatch = errors.New("validator: Es müssen mindestens eine *http.Request und Regeln für die Validate-Methode übergeben werden")
	errInvalidArgument      = errors.New("validator: Ungültige Anzahl an Argumenten übergeben")
	errRequirePtr           = errors.New("validator: Die Datenstruktur muss als Zeiger übergeben werden")
	errRequireData          = errors.New("validator: Bitte geben Sie eine nicht-leere Datenstruktur für die ValidateStruct-Methode an")
	errRequestNotAccepted   = errors.New("validator: Die ValidateStruct-Methode akzeptiert keine *http.Request-Objekte")
)
