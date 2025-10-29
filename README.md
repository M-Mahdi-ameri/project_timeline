# project_timeline


## 👨‍💻 نویسنده
**Mahdi Ameri**


# Timeline Service (Go + Fiber + Redis + MySQL)

## Stack Used

- **Go Fiber** → برای ساخت HTTP Server
- **MySQL + GORM** → برای ذخیره داده‌ها
- **Redis (ZSET)** → برای تایم‌لاین و pagination
- **JWT + bcrypt** → برای احراز هویت امن
- **go-playground/validator** → برای بررسی ورودی‌ها
- **Fan-out Worker** → برای اضافه‌کردن پست‌ها به تایم‌لاین فالوئرها

## ساختار پروژه

```
timeline/
├─ cmd/app/main.go
├─ internal/
│  ├─ config/       # اتصال MySQL و Redis
│  ├─ domain/       # مدل‌ها و پورت‌ها
│  ├─ repository/   # پیاده‌سازی دیتابیس با GORM
│  ├─ handlers/     # کنترلرها (Auth, Posts, Follows, Timeline)
│  ├─ service/      # Worker Fan-out
│  └─ utils/        # JWT و Middleware
├─ .env
└─ README.md
```


```
```
## اجرای پروژه

1️- اجرای MySQL و Redis به صورت محلی (مثلاً با XAMPP و Redis Server):

- MySQL روی پورت 3306
- Redis روی پورت 6379

2️-اجرای پروژه:
```bash
go run cmd/app/main.go
```

خروجی موفق:
```
Connected to MySQL successfully
Connected to Redis successfully
 Server running on port 3000
```
---

## 🪪 Endpointها

| Type | Endpoint | Description | Protected |
|------|-----------|--------------|------------|
| POST | `/register` | ثبت‌نام کاربر | no |
| POST | `/login` | ورود و گرفتن JWT | no |
| POST | `/posts` | ایجاد پست جدید | yes |
| DELETE | `/posts/:id` | حذف پست | yes |
| GET | `/posts/:id` | گرفتن پست با ID | yes |
| GET | `/posts/author/:id` | پست‌های یک نویسنده | yes |
| GET | `/users/:id` | گرفتن اطلاعات کاربر | yes |
| DELETE | `/users/:id` | حذف کاربر | yes |
| POST | `/follow/:id` | دنبال کردن کاربر | yes |
| DELETE | `/unfollow/:id` | آنفالو | yes |
| GET | `/followers/:id` | لیست فالوئرها | yes |
| GET | `/following/:id` | لیست فالووینگ | yes |
| GET | `/timeline` | تایم‌لاین شخصی با pagination | yes |

---

## Pagination Scroll Infinite

پروژه از **Infinite Scroll Pagination** استفاده می‌کند  
(کاربر هر بار پست‌های قدیمی‌تر را می‌گیرد بدون صفحه‌بندی سنتی).

```
GET /timeline?before=<timestamp>&limit=20
```

با استفاده از Redis ZSET:
- **Score:** زمان پست (UnixMilli)
- **Member:** شناسه پست

---

## Fan-out Worker

وقتی کاربر پستی ایجاد می‌کند:
1. پست در MySQL ذخیره می‌شود.
2. Worker لیست فالوئرهای او را می‌گیرد.
3. پست با timestamp به ZSET تایم‌لاین هر فالوئر اضافه می‌شود.

---

## Auth (JWT)

- کاربران جدید با `/register` ساخته می‌شوند.
- رمزها با bcrypt هش می‌شوند.
- در `/login` توکن JWT داده می‌شود.
- سایر مسیرها فقط با Header زیر قابل دسترسی هستند:

```
Authorization: Bearer <your_token_here>
```
