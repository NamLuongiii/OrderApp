export default function ProductsPage() {
  return (
    <div className="min-h-screen p-8 bg-zinc-50 dark:bg-black">
      <div className="max-w-7xl mx-auto">
        <h1 className="text-3xl font-bold mb-6 text-black dark:text-white">
          Sản phẩm
        </h1>
        
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {[1, 2, 3].map((item) => (
            <div key={item} className="bg-white dark:bg-zinc-900 rounded-lg shadow p-6">
              <h3 className="text-xl font-semibold mb-2 text-black dark:text-white">
                Sản phẩm {item}
              </h3>
              <p className="text-zinc-600 dark:text-zinc-400">
                Mô tả sản phẩm mẫu
              </p>
              <div className="mt-4">
                <span className="text-2xl font-bold text-black dark:text-white">
                  {(item * 100000).toLocaleString('vi-VN')}₫
                </span>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
