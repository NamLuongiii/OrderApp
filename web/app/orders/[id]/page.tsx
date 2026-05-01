export default function OrderDetailPage({ params }: { params: { id: string } }) {
  return (
    <div className="min-h-screen p-8 bg-zinc-50 dark:bg-black">
      <div className="max-w-4xl mx-auto">
        <h1 className="text-3xl font-bold mb-6 text-black dark:text-white">
          Chi tiết đơn hàng #{params.id}
        </h1>
        
        <div className="bg-white dark:bg-zinc-900 rounded-lg shadow p-6">
          <div className="space-y-4">
            <div>
              <span className="font-semibold text-black dark:text-white">Mã đơn hàng:</span>
              <span className="ml-2 text-zinc-600 dark:text-zinc-400">{params.id}</span>
            </div>
            <div>
              <span className="font-semibold text-black dark:text-white">Trạng thái:</span>
              <span className="ml-2 text-green-600">Đang xử lý</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
