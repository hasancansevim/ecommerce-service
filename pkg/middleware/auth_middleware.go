package middleware

/*
func AuthMiddleware(authService _interface.AuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, _errors.ErrInvalidToken)
			}
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, _errors.ErrInvalidToken)
			}
			token := tokenParts[1]
			userId, err := jwt.ValidateToken(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, _errors.ErrInvalidToken)
			}
			c.Set("userId", userId)
			return next(c)
		}
	}
}
*/
