import React, { useContext } from "react";
import { Link, useParams, useNavigate } from "react-router-dom";
import { Button, Toast } from "flowbite-react";
import { useQuery, useMutation } from "react-query";

import { API } from "@/lib/api";
import Layout from "@/layouts/Default";
import { toCurrency } from "@/lib/currency";
import { AppContext } from "@/context/AppContext";

const Detail = () => {
	const { id } = useParams();
	const navigate = useNavigate();
	const [state, dispatch] = useContext(AppContext);

	let { data: myCoffee, isLoading: myCoffeeLoading } = useQuery(
		"getBookingCache",
		async () => {
			try {
				const response = await API.get(`/product/${id}`);
				return response.data.data;
			} catch (e) {
				console.log(e);
			}
		}
	);

	const handleAddCart = useMutation(async () => {
		try {
			const response = await API.post(`/order`, {
				product_id: parseInt(id),
			});
			if (response.data.status === "success") {
				navigate("/cart");
			}
		} catch (e) {
			console.log(e.response.data.message);
		}
	});
	return (
		<Layout className='max-w-screen-lg mx-auto'>
			<section className='h-full mt-6 lg:mt-32 lg:px-6'>
				{myCoffeeLoading ? (
					<div className='flex h-[50vh] w-full items-center justify-center'>
						<b className='text-3xl'>Loading Your Product</b>
					</div>
				) : (
					<div className='flex flex-col lg:flex-row my-auto'>
						<div className='lg:w-[46rem]'>
							<img
								src={myCoffee?.image}
								className={"object-cover object-center h-full w-full"}
								alt=''
							/>
						</div>
						<div className='w-full p-6 md:p-12'>
							<h1 className='text-5xl font-bold text-coffee-400'>
								{myCoffee?.name}
							</h1>
							<p className='mt-3 text-coffee-300'>Stock {myCoffee?.stock}</p>
							<p className='text-justify mt-8'>
								Lorem ipsum dolor sit amet consectetur adipisicing elit. Iusto
								accusamus atque suscipit placeat architecto! Asperiores harum
								vero est deserunt itaque magnam omnis, voluptate, maxime quidem,
								quod ratione accusamus illo minus. Lorem, ipsum dolor sit amet
								consectetur adipisicing elit. Ab nostrum odio delectus
								excepturi? Tenetur aspernatur inventore ex, error modi tempora
								eveniet expedita corporis laudantium autem, totam optio rerum
								dicta eos!
							</p>
							<p className='text-right text-3xl mt-8 mb-12 font-bold text-coffee-300'>
								{toCurrency(myCoffee?.price)}
							</p>

							<Button
								type='button'
								size={"lg"}
								className='bg-coffee-400 hover:bg-transparent border-2 !border-coffee-400 text-white hover:text-coffee-400 px-4 py-0 font-semibold transition-all !w-full md:w-auto'
								onClick={() => {
									if (state !== "" || state !== undefined) {
										if (myCoffee?.stock !== 0) {
											handleAddCart.mutate();
										} else {
											Toast.fire({
												icon: "error",
												title: "Product out of stock",
											});
										}
									} else {
										Toast.fire({
											icon: "error",
											title: "You must be logged in to continue !",
										});
									}
								}}
							>
								Add Cart
							</Button>
						</div>
					</div>
				)}
			</section>
		</Layout>
	);
};

export { Detail };
