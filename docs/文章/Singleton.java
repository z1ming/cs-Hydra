public class Singleton {
    private static Singleton instence = null;
    private Singleton() {
        
    }
    
    public static Singleton getInstance() {
        if (instence == null) {
            synchronized (Singleton.class) {
                if (instence == null) {
                    instence = new Singleton();
                }
            }
        }
        return instence;
    }

    
}
