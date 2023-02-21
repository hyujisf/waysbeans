import React, { useContext } from "react";
import { Link, useParams, useNavigate } from "react-router-dom";
import { Button } from "flowbite-react";
import { useQuery, useMutation } from "react-query";

import { API } from "@/lib/api";
import Layout from "@/layouts/Default";
import { toCurrency } from "@/lib/currency";
import { AppContext } from "@/context/AppContext";
import Toast from "@/lib/alert";

const Detail = () => {
	const { id } = useParams();
	const navigate = useNavigate();
	const [state, dispatch] = useContext(AppContext);

	let { data: product, isLoading: productIsLoading } = useQuery(
		"productCache",
		async () => {
			try {
				const response = await API.get(`/product/${id}`);
				return response.data.data;
			} catch (e) {
				console.log(e);
			}
		}
	);

	const handleSubmit = useMutation(async (e) => {
		try {
			e.preventDefault();

			const config = {
				headers: {
					"Content-type": "application/json",
				},
			};
			const body = JSON.stringify({
				product_id: parseInt(id),
				qty: 1,
				sub_total: product?.price,
			});
			await API.post("/order", body, config);
			navigate("/cart");
		} catch (error) {
			console.log(error);
		}
	});

	const title = "Product";
	document.title = "WaysBeans | " + title;

	return (
		<Layout className='max-w-screen-lg mx-auto'>
			<section className='h-full mt-6 lg:mt-32 lg:px-6'>
				{productIsLoading ? (
					<div className='flex h-[50vh] w-full items-center justify-center'>
						<b className='text-3xl'>Loading Your Product</b>
					</div>
				) : (
					<div className='flex flex-col lg:flex-row my-auto'>
						<div className='lg:w-[46rem]'>
							<img
								src={product?.image}
								className={"object-cover object-center h-full w-full"}
								alt=''
							/>
						</div>
						<div className='w-full p-6 md:p-12'>
							<h1 className='text-5xl font-bold text-coffee-400'>
								{product?.name}
							</h1>
							<p className='mt-3 text-coffee-300'>Stock {product?.stock}</p>
							<p className='text-justify mt-8 h-52 overflow-hidden overflow-y-auto'>
								{product?.description}
							</p>
							<div className='text-right text-3xl mt-8 mb-12 font-bold text-coffee-300'>
								{toCurrency(product?.price)}
							</div>
							<Link to='/cart'>
								<Button
									type='button'
									size={"lg"}
									className='bg-coffee-400 hover:bg-transparent border-2 !border-coffee-400 text-white hover:text-coffee-400 px-4 py-0 font-semibold transition-all !w-full md:w-auto'
									onClick={(e) => handleSubmit.mutate(e)}
								>
									Add Cart
								</Button>
							</Link>
						</div>
					</div>
				)}
			</section>
		</Layout>
	);
};

export { Detail };
