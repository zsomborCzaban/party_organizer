package email

func ParseForgotPasswordEmailBody(frontendLink string) string {
	return `
<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333; margin: 0; padding: 0;">
    <div style="max-width: 600px; margin: 0 auto; padding: 20px;">
        <h2 style="color: #2d3748;">Password Reset Request</h2>
        <p style="color: #2d3748;">Hello,</p>
        <p style="color: #2d3748;">We received a request to reset your password. If you didn't make this request, you can safely ignore this email.</p>
        <p style="color: #2d3748;">To reset your password, please click the button below:</p>
        
        <div style="text-align: center; margin: 20px 0;">
            <a href="` + frontendLink + `" 
               target="_blank"
               style="display: inline-block; padding: 12px 24px; background-color: #4299e1; color: white; text-decoration: none; border-radius: 8px; font-weight: bold;">
                Reset Password
            </a>
        </div>
        
        <!-- Plain link fallback -->
        <p>Or copy and paste this link into your browser:<br>
        <a href="` + frontendLink + `" target="_blank" style="color: #4299e1; word-break: break-all;">` + frontendLink + `</a></p>
        
        <div style="margin-top: 30px; font-size: 12px; color: #666;">
            <p>Best regards,<br>Your Party Organizer Team</p>
            <p>This is an automated message, please do not reply to this email.</p>
        </div>
    </div>
</body>
</html>
`
}

func ParseConfirmEmailEmailBody(frontendLink string) string {
	return `
<!DOCTYPE html>
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #2d3748; margin: 0; padding: 0;">
    <div style="max-width: 600px; margin: 0 auto; color: #2d3748; padding: 20px;">
        <h2 style="color: #2d3748;">Confirm Your Email</h2>
        <p style="color: #2d3748;">Hello,</p>
        <p style="color: #2d3748;">Thank you for registering with Party Organizer! To complete your registration, please confirm your email address by clicking the button below:</p>
        <p style="text-align: center;">
            <a href="` + frontendLink + `" style="display: inline-block; padding: 12px 24px; background-color: #4299e1; color: white; text-decoration: none; border-radius: 8px; font-weight: bold;">Confirm Email</a>
        </p>
        <p>Or click this link: <a href="` + frontendLink + `" style="color: #4299e1; word-break: break-all;">` + frontendLink + `</a></p>
        <div style="margin-top: 30px; font-size: 12px; color: #666;">
            <p>Best regards,<br>Your Party Organizer Team</p>
            <p>This is an automated message, please do not reply to this email.</p>
        </div>
    </div>
</body>
</html>
`
}
