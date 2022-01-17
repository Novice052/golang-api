/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	router := NewRouter()
	handler := cors.Default().Handler(router)
	portNumber := "9000"
	fmt.Println("Server listening on port ", portNumber)
	http.ListenAndServe(":"+portNumber, handler)
}
