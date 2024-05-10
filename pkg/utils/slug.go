package utils

import (
    "encoding/base64"
)

// GenerateSlug generates a unique slug of a maximum length of 6 characters
// based on the provided auto-incrementing ID.
func GenerateSlug(id int64) string {
    // Convert the ID to a byte array
    idBytes := make([]byte, 8)
    for i := 7; i >= 0; i-- {
        idBytes[i] = byte(id & 0xff)
        id >>= 8
    }

    // Encode the byte array using base64
    base64Slug := base64.URLEncoding.EncodeToString(idBytes)

    // Truncate the base64 slug to 6 characters
    if len(base64Slug) > 6 {
        base64Slug = base64Slug[:6]
    }

    return base64Slug
}