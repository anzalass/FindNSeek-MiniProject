# Find N Seek

Find N Seek memiliki arti mencari dan menemukan, adalah API Server pencarian barang atau lainya yang hilang, program ini ditulis menggunakan Go sebagai bahasa utama, echo sebagai framework, gorm sebagai orm, dan mysql sebagai database nya, api ini diharapkan bisa memudahkan siapapun yang sedang mencari sesuatu yang hilang

### Gambaran Singkat mengenai API Find N Seek
  1. Setiap user dapat mempublish barang yang dicari
  2. Setiap user juga dapat mengajukan barang yang ditemukan ke user yang mencari
  3. Setiap ada pengajuan, user yang mencari akan mendapatkan notifikasi berupa gambar, yang dikirim melalui email
  4. Ketika benar barang tersebut adalah miliknya maka user yang mencari bisa melakukan claim, bahwa benar barang tersebut adalah miliknya
  5. Kemudian user yang mengajukan akan dapat notifikasi berisi kontak user yang mencari, yang dikirim melalui email

[ERD](https://lucid.app/lucidchart/7d2520c7-57c8-4b63-a8f4-83c51fed858f/edit?viewport_loc=-2923%2C-2859%2C11082%2C5042%2C0_0&invitationId=inv_f1fe26c2-4691-402a-b714-395326c1e371 "Find N Seek ERD").

![Find N Seek)](https://github.com/anzalass/FindNSeek-MiniProject/assets/109114576/a0e69fad-a2b5-4519-b707-9b1da94d9f45)

[API Documentation](https://documenter.getpostman.com/view/25780742/2s9YRGy98F "Find N Seek Documentation").
