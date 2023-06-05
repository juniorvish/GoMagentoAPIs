import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;

public class Main {

    public static void main(String[] args) {
        ScheduledExecutorService executor = Executors.newScheduledThreadPool(3);

        SyncProducts syncProducts = new SyncProducts();
        SyncCustomers syncCustomers = new SyncCustomers();
        SyncOrders syncOrders = new SyncOrders();

        executor.scheduleAtFixedRate(syncProducts::sync, 0, 30, TimeUnit.MINUTES);
        executor.scheduleAtFixedRate(syncCustomers::sync, 0, 30, TimeUnit.MINUTES);
        executor.scheduleAtFixedRate(syncOrders::sync, 0, 30, TimeUnit.MINUTES);
    }
}