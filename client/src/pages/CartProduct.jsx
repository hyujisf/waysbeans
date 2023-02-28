import React, { useState, useEffect } from "react";
import { IoMdTrash } from "react-icons/io";
import { Button } from "flowbite-react";
import { useNavigate } from "react-router-dom";
import { useMutation, useQuery } from "react-query";

import { API } from "@/lib/api";
import Layout from "@/layouts/Default";
import { toCurrency } from "@/lib/currency";

const CartProduct = () => {
	const navigate = useNavigate;

	//cart
	let { data: cart, refetch } = useQuery("cartsCache", async () => {
		const response = await API.get("/orders-id");
		return response.data.data;
	});

	//subtotal
	let resultTotal = cart?.reduce((x, y) => {
		return x + y.qty * y.subtotal;
	}, 0);

	//qty
	let resultQty = cart?.reduce((x, y) => {
		return x + y.qty;
	}, 0);

	//remove
	let handleDelete = async (id) => {
		await API.delete(`/order/` + id);
		refetch();
	};

	//update
	const increaseCart = async (idProduct) => {
		try {
			const result = cart.find(({ id }) => id === idProduct);

			const config = {
				headers: {
					"Content-type": "application/json",
				},
			};

			const body = JSON.stringify({
				qty: result.qty + 1,
			});

			await API.patch("/order/" + idProduct, body, config);
			refetch();
		} catch (error) {}
	};

	const decreaseCart = async (idProduct) => {
		try {
			const result = cart.find(({ id }) => id === idProduct);

			const config = {
				headers: {
					"Content-type": "application/json",
				},
			};

			const body = JSON.stringify({
				qty: result.qty - 1,
			});

			await API.patch("/order/" + idProduct, body, config);
			refetch();
		} catch (error) {
			console.log(error);
		}
	};

	// pay midtrans
	useEffect(() => {
		//change this to the script source you want to load, for example this is snap.js sandbox env
		const midtransScriptUrl = "https://app.sandbox.midtrans.com/snap/snap.js";
		//change this according to your client-key
		// const myMidtransClientKey = "Client key here ...";
		const myMidtransClientKey = import.meta.env.VITE_MIDTRANS_CLIENT_KEY;

		let scriptTag = document.createElement("script");
		scriptTag.src = midtransScriptUrl;
		// optional if you want to set script attribute
		// for example snap.js have data-client-key attribute
		scriptTag.setAttribute("data-client-key", myMidtransClientKey);

		document.body.appendChild(scriptTag);
		return () => {
			document.body.removeChild(scriptTag);
		};
	}, []);

	// handlebuy

	const form = {
		total: resultTotal,
	};
	const handleSubmit = useMutation(async (e) => {
		const config = {
			headers: {
				"Content-type": "application/json",
			},
		};
		const body = JSON.stringify(form);
		const response = await API.post("/transaction", body, config);
		const token = response.data.data.token;

		window.snap.pay(token, {
			onSuccess: function (result) {
				console.log(result);
				navigate("/profile");
			},
			onPending: function (result) {
				console.log(result);
				navigate("/profile");
			},
			onError: function (result) {
				console.log(result);
			},
			onClose: function () {
				alert("you closed the popup without finishing the payment");
			},
		});
		await API.patch("/order", body, config);
	});

	const title = "Cart";
	document.title = "WaysBeans | " + title;
	return (
		<Layout className='max-w-screen-lg mx-auto'>
			<section className='mt-32'>
				<h1 className='text-3xl font-semibold'>My Cart</h1>
				{cart?.length != 0 ? (
					cart?.map((item, k) => (
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
											<h4 className='mb-3 font-semibold'>
												{item?.product.name}
											</h4>
											<div>
												<button onClick={() => decreaseCart(item.id)}>-</button>
												<input
													type='text'
													value={item?.qty}
													className='w-10 h-6 p-0 text-center mx-2 bg-coffee-100 rounded-lg border-0'
												/>
												<button onClick={() => increaseCart(item.id)}>+</button>
											</div>
										</td>
										<td className='text-right'>
											<p className='mb-3'>Rp.300.900</p>
											<button onClick={() => handleDelete(item.id)}>
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
										<td className='text-right'>
											{toCurrency(item?.product?.price)}
										</td>
									</tr>
									<tr>
										<td>Qty</td>
										<td className='text-right'>{item?.qty}</td>
									</tr>
									<tr className='border-t border-x-0 border-t-coffee-400 text-coffee-250 font-semibold'>
										<td>Total</td>
										<td className='text-right'>
											{toCurrency(item?.qty * item?.product?.price)}
										</td>
									</tr>
								</table>
								<div className='flex justify-end'>
									<Button
										type='button'
										onClick={(e) => handleSubmit.mutate(e)}
										className='bg-coffee-400 hover:bg-transparent border-2 !border-coffee-400 text-white hover:text-coffee-400 px-20 py-0 font-semibold transition-all md:w-auto text-right'
									>
										Pay
									</Button>
								</div>
							</div>
						</div>
					))
				) : (
					<div className='text-center'>
						<img
							src={imageTrans}
							className='img-fluid'
							style={{ width: "40%" }}
						/>
						<div className='text-primer fw-bold'>
							No data Cart, let's Shopping
						</div>
					</div>
				)}
			</section>
		</Layout>
	);
};

export { CartProduct };
