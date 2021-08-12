

import java.security.Key;
import java.util.Base64;
import javax.crypto.Cipher;
import javax.crypto.spec.SecretKeySpec;

public class App {
    public static void main(String[] args) throws Exception {

        try {
            // Cria a key e cifra
            Key aesKey = new SecretKeySpec("ozMjQ1MjM1MzJnaw".getBytes(), "AES");
            Cipher cipher = Cipher.getInstance("AES");
            // criptografa o texto
            cipher.init(Cipher.ENCRYPT_MODE, aesKey);
            byte[] encrypted = cipher.doFinal("Senha descriptografada".getBytes());            
           
            System.out.println("----------- codcripto - " + encrypted);

            String base64 = Base64.getEncoder().encodeToString(encrypted);
           
            System.out.println("----------- codf - " + base64);
        }
        catch(Exception e) {
            e.printStackTrace();
        }
       
        try {
            // Cria a key e cifra
            Key aesKey = new SecretKeySpec("ozMjQ1MjM1MzJnaw".getBytes(), "AES");
            Cipher cipher = Cipher.getInstance("AES");
            // descriptografa o texto
            cipher.init(Cipher.DECRYPT_MODE, aesKey);
            byte[] encrypted =  Base64.getDecoder().decode("dzn/e0hiX/pog49BZxMhqkScBuhzdpn5zZprtmwizJo=".getBytes());
            String decrypted = new String(cipher.doFinal(encrypted));
           
            System.out.println("----------- decod - " + decrypted);
        }
        catch(Exception e) {
            e.printStackTrace();
        }

        System.out.println("Hello, World!");
    }
}