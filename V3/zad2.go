// Transakcijski blok

// Deklarairaj strukture
// Transaction: from (string), to (string), amount (int)
// Block: Number (int), time (Time), slice struktura Transaction
// Kreiraj funkciju koja služi kao konstruktor Block-a i prima parametar number
// funkcija u novo kreirani blok upisuje trenutno vrijeme
// inicijalizira Transaction slice
// Block ima metodu AddTxs kroz koju se može dodati jedna ili više struktura Transaction u popis transakcija bloka
// Block nadjačava metodu String() u kojoj vraća string s informacijama o bloku i transakcijama
// ![Block class diagram](imgs/Block class diagram.png)