Git Flow diperkenalkan oleh Vincent Driessen pada tahun 2010.
Pada dasarnya, Git Flow adalah sebuah alur kerja (workflow) dalam sebuah pengembangan software yang memanfaatkan git untuk mengatur proses bagaimana penambahan fitur, pembuatan dokumentasi, rilis versi, dan perbaikan bug dibuat secara terstruktur di dalam sebuah repositori.


Git Flow sangat cocok digunakan saat membuat software yang membutuhkan banyak tim dalam proses pengembangannya, karena bisa membantu menjaga keteraturan dan kejelasan alur kerja. Dengan Git Flow, pekerjaan tim bisa dibagi ke dalam beberapa branch yang punya peran masing-masing, seperti develop, feature, release, dan hotfix. Semua cabang/branch tadi dipisahkan dari branch utama (main) supaya proses pengembangan tidak langsung memengaruhi versi yang akan menjadi production.
Branch-branch tersebut nantinya akan digabung ke branch main setelah melewati berbagai pengecekan dan standar kualitas perusahaan. Hal ini membantu mengurangi risiko konflik dan menjaga stabilitas, sekaligus meningkatkan efisiensi kolaborasi antar anggota tim.


1. Branch Main
   Dalam Git Flow, branch main dibuat sejak awal proyek dimulai dan digunakan sebagai tempat menyimpan kode final yang siap dirilis ke produksi, pengembangan tidak langsung dilakukan di     branch ini, jadi biasanya perubahan yang masuk ke dalam branch main berasal dari branch/cabang lain
2. Branch Develop
   Branch develope adalah Tempat utama yang digunakan untuk menggabungkan semua fitur yang sedang dikembangkan dalam branch feature, sebelum nantinya di lanjutkan ke branch main
3. Branch Feature
   branch yang digunakan untuk mengembangkan fitur baru, branch ini dibuat dari turunan branch develop dan akan digabung ke branch develop kembali, biasaya penamaan branch feature           berdasarkan fitur yang sedang dibuat, seperti feature/login-page, feature/register-page, dan sebagainya
4. Branch Release
   Branch release dibuat ketika proyek sudah mendekati waktu rilis dari software, biasanya hanya berisi perbaikan kecil, penyesuaian akhir, atau penambahan dokumentasi. Setelah siap,        branch ini digabung    ke main dan juga ke develop agar semua update tetap sinkron ke depannya.
5. Branch Hotfix
   Digunakan untuk memperbaiki bug penting atau error darurat yang baru ditemukan saat software sudah dalam versi produksi, branch ini diturunkan langsung dari branch main karena harus      memperbaiki versi yang sudah produksi, dalam proses pengerjaannya harus dikerjakan dengan cepat dan tepat supaya bug tidak berdampak lebih jauh pada pengguna. Setelah masalah             diperbaiki, kode harus digabungkan kembali ke main dan develop agar perbaikan tetap ada di rilisan berikutnya.




![image](https://github.com/user-attachments/assets/014bd6f7-ccd8-4bcb-a19b-ad5c3c988d9c)

PENJELASAN ALUR GIT FLOW DARI DIAGRAM DI ATAS:
1. Main Branch
- Ditandai dengan garis hijau.
- Menyimpan versi yang stabil dan siap produksi.
- Titik pertama adalah rilis v1.0, artinya versi awal aplikasi sudah dirilis.

2. Develop Branch
- Ditandai dengan garis ungu.
- Berasal(turunan) dari branch main, dan digunakan untuk menggabungkan fitur-fitur yang dikembangkan sebelum akhirnya dirilis.
- Setelah v1.0 dirilis, pengembangan berlanjut di develop.

3. Feature Branch
- Ditandai dengan garis merah.
- Feature 1 dan Feature 2 dibuat dari develop untuk mengembangkan fitur baru secara terpisah.
- Setelah selesai dikembangkan, kedua fitur ini digabungkan kembali ke develop.

4. Release Branch (Release v2.0)
- Ketika kumpulan fitur di develop dianggap siap untuk dirilis, dibuatlah branch release/2.0 dari develop.
- Di sinilah dilakukan penyempurnaan kecil sebelum rilis final.
- Setelah jalan dengan stabil, branch ini digabungkan ke:
  branch main → menghasilkan rilis v2.0
  branch develop → agar perubahan tetap sinkron

5. Hotfix Branch (Hotfix for v2.0)
- Setelah v2.0 dirilis, ternyata ditemukan bug.
- Maka dibuatlah branch hotfix langsung dari main.
- Setelah bug diperbaiki:
  Hasilnya digabung ke branch main → menghasilkan v2.1
  Juga digabung ke branch develop agar perbaikan bug tetap tercatat


Kesimpulannya Git Flow adalah strategi manajemen branch yang membantu developer dalam mengatur alur kerja secara terstruktur dan terorganisir, terutama untuk proyek dengan rilis berkala dan tim yang besar. Dengan membagi proses pengembangan ke dalam beberapa branch, sehingga setiap tahapan dari penambahan fitur, persiapan rilis, hingga perbaikan bug bisa dikerjakan secara terpisah tanpa saling mengganggu.
Git Flow menjaga stabilitas kode produksi (main), sekaligus memberikan ruang untuk inovasi (feature) dan perbaikan cepat (hotfix) semuanya tetap sinkron di branch develop.
