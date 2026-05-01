export default function DashboardPage() {
  return (
    <div className="min-h-screen p-8 bg-zinc-50 dark:bg-black">
      <div className="max-w-7xl mx-auto">
        <h1 className="text-3xl font-bold mb-6 text-black dark:text-white">
          Dashboard
        </h1>
        
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
          <div className="bg-white dark:bg-zinc-900 rounded-lg shadow p-6">
            <h3 className="text-sm font-medium text-zinc-600 dark:text-zinc-400">
              Tổng đơn hàng
            </h3>
            <p className="text-3xl font-bold mt-2 text-black dark:text-white">245</p>
          </div>
          
          <div className="bg-white dark:bg-zinc-900 rounded-lg shadow p-6">
            <h3 className="text-sm font-medium text-zinc-600 dark:text-zinc-400">
              Doanh thu
            </h3>
            <p className="text-3xl font-bold mt-2 text-black dark:text-white">
              125M₫
            </p>
          </div>
          
          <div className="bg-white dark:bg-zinc-900 rounded-lg shadow p-6">
            <h3 className="text-sm font-medium text-zinc-600 dark:text-zinc-400">
              Khách hàng
            </h3>
            <p className="text-3xl font-bold mt-2 text-black dark:text-white">89</p>
          </div>
        </div>
        
        <div className="bg-white dark:bg-zinc-900 rounded-lg shadow p-6">
          <h2 className="text-xl font-semibold mb-4 text-black dark:text-white">
            Hoạt động gần đây
          </h2>
          <p className="text-zinc-600 dark:text-zinc-400">
            Danh sách hoạt động sẽ hiển thị ở đây
          </p>
        </div>
      </div>
    </div>
  );
}
