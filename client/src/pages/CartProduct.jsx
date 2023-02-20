import React from "react";
import { IoMdTrash } from "react-icons/io";
import { Button } from "flowbite-react";

import Layout from "@/layouts/Default";

const CartProduct = () => {
	return (
		<Layout className='max-w-screen-lg mx-auto'>
			<section className='mt-32'>
				<h1 className='text-3xl font-semibold'>My Cart</h1>
				<div className='flex gap-4 pt-8 flex-col lg:flex-row'>
					<div className='lg:w-3/5'>
						<h4 className='text-lg mb-2'>Review Your Order</h4>

						<table className='w-full border-t border-x-0 border-t-coffee-400'>
							<tr className='border-b border-x-0 border-b-coffee-400 px-2 py-4 gap-4 flex w-full'>
								<td className='w-[4.5rem]'>
									<div className='w-full'>
										<img
											src={"/uploads/product1.png"}
											className='h-full'
											alt=''
										/>
									</div>
								</td>
								<td className='w-full'>
									<h4 className='mb-3 font-semibold'>GUETEMALA Beans</h4>
									<div>
										<button>-</button>
										<input
											type='text'
											value={2}
											className='w-10 h-6 p-0 text-center mx-2 bg-coffee-100 rounded-lg border-0'
										/>
										<button>+</button>
									</div>
								</td>
								<td className='text-right'>
									<p className='mb-3'>Rp.300.900</p>
									<button>
										<IoMdTrash fontSize={20} />
									</button>
								</td>
							</tr>
						</table>
					</div>
					<div className='lg:w-2/5'>
						<table className='mt-9 w-full border-t border-x-0 border-t-coffee-400 text-coffee-300'>
							<tr>
								<td>Subtotal</td>
								<td className='text-right'>1961</td>
							</tr>
							<tr>
								<td>Qty</td>
								<td className='text-right'>2</td>
							</tr>
							<tr className='border-t border-x-0 border-t-coffee-400 text-coffee-250 font-semibold'>
								<td>Total</td>
								<td className='text-right'>300.900</td>
							</tr>
						</table>
						<div className='flex justify-end'>
							<Button
								type='button'
								className='bg-coffee-400 hover:bg-transparent border-2 !border-coffee-400 text-white hover:text-coffee-400 px-20 py-0 font-semibold transition-all md:w-auto text-right'
							>
								Pay
							</Button>
						</div>
					</div>
				</div>
			</section>
		</Layout>
	);
};

export { CartProduct };
