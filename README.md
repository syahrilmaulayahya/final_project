# final_project (ecommerce untuk penjualan batik)
Project ini bertujuan untuk membuat E-Commerce khusus untuk menjual berbagai macam produk batik, seperti kain batik, kemeja, kebaya, gamis, dan lain-lain, dengan fitur - fitur sederhana seperti register user, login, upload produk, pembelian dan lain-lain.

# MVP(Minimum Viable Product) #


* Register, login dan menampilkan informasi user
* Menampilkan Produk
* Menambahkan rating dan review produk 
* Pencarian produk
* Filter produk berdasarkan kategori
* Masukkan keranjang
* Checkout
* Pembayaran
* Pengiriman dan menampilkan informasinya
* Menampilkan informasi tentang produk yang dibeli sebelumnya
* Pembatalan dan pengembalian produk
* Upload dan update produk

# Link ERD #
[Link ERD](https://drive.google.com/file/d/1W9CN2bIQPpx7rhHZAlgpzt1iT0PkSBaJ/view?usp=sharing)

# Teknologi dan Framework yang digunakan
* echo labstack
* gorm
* jwt
* aws
* mysql
* docker
* viper
* git dan github
* mockery

# List API dan Penjelasannya #
## user ##
* users/logins : login user
* users/details/:id : mendapatkan detail user berdasarkan id
* users/registers : register user
* users/reviews : menambahkan review dan rating suatu produk

## admin ##
* admins/registers : register admin
* admins/logins : login admin

## product ##
* product : menampilkan semua produk
* products/details/:id : menampilkan detail produk berdasarkan id
* products/searches : mencari produk berdasarkan nama
* products/filters : memfilter produk berdasarkan product type id
* products/uploads : mengupload produk baru (admin)
* products/updates : mengupdate produk yang sudah ada (admin)
* products/uploadtypes : mengupload type produk baru (admin)
* products/uploadsizes : mengupload size baru (admin)
* products/updatesizes : mengupdate size yang sudah ada (admin)
* products/updatestocks : update stock salah satu size (admin)
* products/uploaddescriptions : upload deskripsi produk (admin)
* products/updatedescriptions : mengupdate deskripsi produk yang sudah ada (admin)

## transactions ##
* transactions/addshoppingcarts : menambahkan produk ke keranjang
* transactions/details : menampilkan detail keranjang belanja
* transactions/details/checkouts : melakukan checkout produk yang sudah ada dalam keranjang
* transactions/details/checkouts/pns : memilih metode pembayaran dan pengiriman (jika tidak memilih, akan digunakan format default)
* transactions/details/checkouts/pns/pay : melakukan pembayaran sesuai dengan nominal yang ada di transactions
* transactions/addpayments : menambahkan pilihan pembayaran (admin)
* transactions/getpayments : menampilkan semua pilihan pembayaran
* transactions/addshipments : menambahakan pilihan pengiriman (admin)
* transactions/getshipments : menampilkan semua pilihan pengiriman
* transactions/transactiondetails : menampilkan detail transaksi setelah melakukan pembayaran
* transactions/transactiondetails/delivered : mengganti status pengiriman (undelivered menjadi delivered)
* transactions/transactiondetails/canceled : membatalkan produk yang sudah dibeli

