export default function Home({ message }: { message: string }) {
  return (
    <div className="min-h-screen bg-gray-50 flex items-center justify-center p-6">
      <div className="max-w-2xl w-full bg-white rounded-lg shadow-lg p-8">
        <h1 className="text-4xl font-bold text-gray-900 mb-4">Welcome to Velocity</h1>
        <p className="text-lg text-gray-600 mb-6">
          Velocity brings Laravel's developer experience to Go without sacrificing performance.
        </p>
        <div className="space-y-4">
          <div className="flex items-start gap-3">
            <span className="flex h-6 w-6 items-center justify-center rounded-full bg-blue-100 text-blue-600 text-sm font-medium">1</span>
            <div>
              <p className="text-gray-700">
                Read the <a href="https://github.com/velocitykode/velocity" target="_blank" rel="noopener noreferrer" className="text-blue-600 hover:text-blue-800 underline">Documentation</a>
              </p>
            </div>
          </div>
          <div className="flex items-start gap-3">
            <span className="flex h-6 w-6 items-center justify-center rounded-full bg-blue-100 text-blue-600 text-sm font-medium">2</span>
            <div>
              <p className="text-gray-700">
                Explore <a href="https://github.com/velocitykode/velocity/tree/main/examples" target="_blank" rel="noopener noreferrer" className="text-blue-600 hover:text-blue-800 underline">Examples</a>
              </p>
            </div>
          </div>
        </div>
        <div className="mt-8 pt-6 border-t border-gray-200">
          <a
            href="https://github.com/velocitykode/velocity"
            target="_blank"
            rel="noopener noreferrer"
            className="inline-block bg-gray-900 text-white px-6 py-2 rounded-md hover:bg-gray-800 transition"
          >
            Star on GitHub
          </a>
        </div>
      </div>
    </div>
  )
}
