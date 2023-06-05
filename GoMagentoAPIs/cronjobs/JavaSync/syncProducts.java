import java.time.LocalDateTime;
import java.util.List;
import java.util.Timer;
import java.util.TimerTask;

public class SyncProducts {

    private static final String MAGENTO_BASE_URL = "https://your-magento-instance.com";
    private static final String DESKERA_BASE_URL = "https://api.deskera.com";
    private static final String AUTH_TOKEN = "your-auth-token";

    public static void main(String[] args) {
        Timer timer = new Timer();
        timer.schedule(new SyncProductsJob(), 0, 1800000); // 30 minutes in milliseconds
    }

    static class SyncProductsJob extends TimerTask {
        @Override
        public void run() {
            syncProducts();
        }
    }

    public static void syncProducts() {
        LocalDateTime lastSyncTime = getLastSyncTime();
        List<Product> newProducts = getNewProductsFromMagento(lastSyncTime);
        for (Product product : newProducts) {
            DeskeraProduct deskeraProduct = mapProductToDeskeraProduct(product);
            createDeskeraProduct(deskeraProduct);
        }
        updateLastSyncTime();
    }

    private static LocalDateTime getLastSyncTime() {
        // Implement logic to get the last sync time from a persistent storage
        return LocalDateTime.now().minusMinutes(30);
    }

    private static List<Product> getNewProductsFromMagento(LocalDateTime lastSyncTime) {
        // Implement logic to call Magento API and get new products created after lastSyncTime
        return null;
    }

    private static DeskeraProduct mapProductToDeskeraProduct(Product product) {
        // Implement logic to map Magento Product to Deskera Product
        return null;
    }

    private static void createDeskeraProduct(DeskeraProduct deskeraProduct) {
        // Implement logic to call Deskera API and create a new product
    }

    private static void updateLastSyncTime() {
        // Implement logic to update the last sync time in a persistent storage
    }
}