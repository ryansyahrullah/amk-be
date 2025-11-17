package mail

// File: pkg/mail/mailer.go
// ---------------------------------------------------------
// Utility untuk mengirim email.
// Nantinya menggunakan library: gopkg.in/gomail.v2
//
// Fungsi utama:
//   - SendMail(to, subject, body)
//
// Digunakan oleh:
//   - auth/service/password_service.go untuk mengirim OTP lupa password.
// Konfigurasi SMTP diambil dari environment (.env):
//   - SMTP_HOST, SMTP_PORT, SMTP_USER, SMTP_PASS, SMTP_FROM.
