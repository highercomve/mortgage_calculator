export function Details({ data }) {
  if (!data) {
    return null
  }
  return (
    <div class="mt-11 overflow-hidden bg-white shadow sm:rounded-lg">
      <div class="px-4 py-5 sm:px-6">
        <h3 class="text-lg font-medium leading-6 text-gray-900">Mortgage information</h3>
        <p class="mt-1 max-w-2xl text-sm text-gray-500">Here are the details of you mortgage.</p>
      </div>
      <div class="border-t border-gray-200">
        <dl>
        <div class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
            <dt class="text-sm font-medium text-gray-500">Schedule</dt>
            <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0 capitalize">{data.schedule}</dd>
          </div>
          <div class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
            <dt class="text-sm font-medium text-gray-500">Number of payments</dt>
            <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">{data.number_of_payments}</dd>
          </div>
          <div class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
            <dt class="text-sm font-medium text-gray-500">Payment by schedule</dt>
            <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">{data.payment}</dd>
          </div>
        </dl>
      </div>
    </div>
  )
}