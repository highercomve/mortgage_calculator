import { useRef, useState } from "react";
import { PostCalculate } from "../../services/calculator";

export function CalculateForm({ onResponse = () => {} }) {
  const formRef = useRef()
  const [loading, setLoading] = useState()

  const onSubmit = async (event) => {
    event.preventDefault();
    setLoading(true)

    const form = formRef.current
    const isValid = form.checkValidity()
    if (!isValid) {
      // TODO: Manage better if the form is invalid for now it will validated by the browser
      return alert("the form is not valid")
    }

    const formData = new FormData(form)
    const data = Object.fromEntries(formData.entries())

    // Convert the type of data to the correct type of data.
    const resp = await PostCalculate({
      price: Number(data.price),
      downpayment: Number(data.downpayment),
      period: Number(data.period),
      anual_interest: Number(data.anual_interest),
      schedule: data.schedule,
    })

    // Send response to on response function
    // Maybe we could process if is an error in order to send to different functions
    onResponse(resp)
    setLoading(false)
  }

  return (
    <form
      className="mt-8 space-y-6"
      action="#"
      method="POST"
      ref={formRef}
      onSubmit={onSubmit}
    >
      <div className="p-4 -space-y-px rounded-md">

        <div>
          <label htmlFor="price" className="my-3 block">Price</label>
          <input
            id="price"
            name="price"
            type="number"
            step="0.01"
            required
            className="relative block w-full rounded-none rounded-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
            placeholder="200000.00"
          />
        </div>
        <div>
          <label htmlFor="downpayment" className="my-3 block">Downpayment</label>
          <input
            id="downpayment"
            name="downpayment"
            type="number"
            step="0.01"
            required
            className="relative block w-full rounded-none rounded-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
            placeholder="20000.00"
          />
        </div>
        <div>
          <label htmlFor="anual_interest" className="my-3 block">Anual Interest</label>
          <input
            id="anual_interest"
            name="anual_interest"
            type="number"
            step="0.1"
            required
            className="relative block w-full rounded-none rounded-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
            placeholder="5"
          />
        </div>
        <div>
          <label htmlFor="schedule" className="my-3 block">Schedule</label>
          <select
            id="schedule"
            name="schedule"
            required
            className="relative block w-full rounded-none rounded-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
            placeholder="Schedule"
          >
            <option value="monthly">Monthly</option>
            <option value="bi-weekly">Bi-Weekly</option>
            <option value="accelerated bi-weekly">Accelerated Bi-Weekly</option>
          </select>
        </div>
        <div>
          <label htmlFor="period" className="my-3 block">Period</label>
          <select
            id="period"
            name="period"
            required
            className="relative block w-full rounded-none rounded-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
            placeholder="Period"
          >
            <option value="5">5</option>
            <option value="10">10</option>
            <option value="15">15</option>
            <option value="20">20</option>
            <option value="25">25</option>
            <option value="30">30</option>
          </select>
        </div>
        <div className="py-6 flex items-center justify-center">
          <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            type="submit"
            disabled={loading}
          >
            Calculate
          </button>
        </div>
      </div>
    </form>
  )
}