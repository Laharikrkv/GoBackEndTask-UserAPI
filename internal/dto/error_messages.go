package dto

import ("fmt")

type ApiError struct {
    Field string `json:"field"`
    Msg   string `json:"msg"`
}

func MsgForTag(tag string) string {
    switch tag {
    case "required":
        return "This field is required"
    case "alphaspace":
        return "Should contain letters and spaces"
    case "min":
        return "Must have atleast two characters"
	case "datetime":
		return "DOB must in yyyy-mm-dd Format"
    default:
        return fmt.Sprintf("Field failed on '%s' validation", tag)
    }
}