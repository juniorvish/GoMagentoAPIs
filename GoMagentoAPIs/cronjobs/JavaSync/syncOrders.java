import java.time.LocalDateTime;
import java.util.List;
import java.util.Timer;
import java.util.TimerTask;

public class SyncOrders {

    private static final String MAGENTO_BASE_URL = "https://your-magento-instance.com";
    private static final String DESKERA_BASE_URL = "https://api.deskera.com";
    private static final String AUTH_TOKEN = "your-auth-token";

    public static void main(String[] args) {
        Timer timer = new Timer();
        timer.schedule(new SyncOrdersJob(), 0, 1800000); // 30 minutes in milliseconds
    }

    static class SyncOrdersJob extends TimerTask {
        @Override
        public void run() {
            syncOrders();
        }

        private void syncOrders() {
            LocalDateTime lastSyncTime = getLastSyncTime();
            List<Order> newOrders = getNewOrders(lastSyncTime);

            for (Order order : newOrders) {
                DeskeraSalesInvoice deskeraInvoice = mapOrderToDeskeraInvoice(order);
                createDeskeraSalesInvoice(deskeraInvoice);
            }

            updateLastSyncTime(LocalDateTime.now());
        }

        private LocalDateTime getLastSyncTime() {
            // Implement logic to get the last sync time from a persistent storage
            return LocalDateTime.now().minusMinutes(30);
        }

        private List<Order> getNewOrders(LocalDateTime lastSyncTime) {
            // Implement logic to call Magento API to get orders created after lastSyncTime
            return null;
        }

        private DeskeraSalesInvoice mapOrderToDeskeraInvoice(Order order) {
            // Implement logic to map Magento Order object to Deskera Sales Invoice object
            return null;
        }

        private void createDeskeraSalesInvoice(DeskeraSalesInvoice invoice) {
            // Implement logic to call Deskera API to create a new sales invoice
        }

        private void updateLastSyncTime(LocalDateTime newSyncTime) {
            // Implement logic to update the last sync time in a persistent storage
        }
    }
}