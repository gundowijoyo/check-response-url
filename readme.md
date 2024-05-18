# Pengecek Kecepatan Internet dengan Go

Proyek ini adalah program sederhana dalam Go untuk mengecek kecepatan internet dengan mengukur waktu respons dari URL yang ditentukan. Program ini menggunakan loop terus-menerus untuk secara berkala mengirim permintaan HTTP GET ke daftar URL dan mencetak waktu responsnya.

## Fitur

- Mengukur waktu respons dari beberapa URL.
- Pemantauan terus-menerus dengan interval yang dapat disesuaikan.
- Sederhana dan mudah dipahami.

## Prasyarat

- [Go](https://golang.org/dl/) terpasang di komputer Anda.

## Penggunaan

1. **Clone repositori:**

    ```sh
    git clone https://github.com/guns-joy/check-response-url.git cd check-response-url ```

2. **Jalankan program:**

    ```sh
    go run interface.go
    ```

3. **Modifikasi URL:**

    Anda dapat memodifikasi daftar URL yang akan diuji dengan mengedit slice `urls` di `main.go`:

    ```go
    urls := []string{
        "https://www.google.com",
        "https://www.cloudflare.com",
        // Tambahkan URL lainnya di sini
    }
    ```

4. **Sesuaikan interval:**

    Interval antara pengecekan diatur menjadi 10 detik. Anda dapat mengubahnya dengan memodifikasi durasi `time.Sleep` di `main.go`:

    ```go
    time.Sleep(10 * time.Second) // Sesuaikan interval di sini
    ```

## Contoh Output
![Alt Text](https://s10.gifyu.com/images/SfTDm.jpg)
