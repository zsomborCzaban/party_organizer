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
        <p>Hello,</p>
        <p>We received a request to reset your password. If you didn't make this request, you can safely ignore this email.</p>
        <p>To reset your password, please click the button below:</p>
        
        <!-- Button with full inline styling -->
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
