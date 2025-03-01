# Internship   Problem Solving Test
# Nama -  Ajri Inda Robby

# Internship - Problem Solving Test SE (BE)

/////////////////////////////////////////////////////////////////////////////////////////////

Penjelasan Kode Soal no 3:

Menggunakan stack []rune{} untuk menyimpan bracket buka.

Memanfaatkan map[rune]rune untuk mencocokkan bracket tutup dengan pasangan yang benar.

Iterasi string:

    Jika karakter adalah bracket buka { [ (, push ke stack.
  
    Jika karakter adalah bracket tutup } ] ):
  
      Jika stack kosong atau bracket tidak cocok, return "NO".
    
      Jika cocok, pop dari stack.
    
Setelah iterasi:

    Jika stack kosong, return "YES", artinya semua bracket telah dipasangkan.
  
    Jika stack masih berisi elemen, return "NO", artinya ada bracket buka yang tidak memiliki pasangan.


