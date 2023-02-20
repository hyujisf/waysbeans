import React, { useContext } from "react";
import { Link, useNavigate } from "react-router-dom";
import { Card } from "flowbite-react";
import { useQuery } from "react-query";

import { API } from "@/lib/api";
import { toCurrency } from "@/lib/currency";
import { AppContext } from "@/context/AppContext";
import Layout from "@/layouts/Default";

const Home = () => {
	const navigate = useNavigate();
	const [state, dispatch] = useContext(AppContext);

	const { data: productData, isLoading: productDataIsLoading } = useQuery(
		"productDataCache",
		async (e) => {
			try {
				const response = await API.get("/products");
				return response.data.data;
			} catch (err) {
				console.log(err);
			}
		}
	);

	document.title = "WaysBeans";
	return (
		<Layout className='max-w-screen-lg mx-auto'>
			<>
				<section className='mt-10'>
					<div className='bg-coffee-200 p-6 lg:p-12 mr-24'>
						<div className='flex flex-col lg:flex-row'>
							<div className='lg:w-8/12'>
								<img src={"/img/Waysbeans-lg.svg"} alt='WaysBeans Logo' />
								<h2 className='text-md sm:text-2xl mt-3 mb-4 md:mb-6'>
									BEST QUALITY COFFEE BEANS
								</h2>
								<p className='text-sm sm:text-lg lg:w-9/12 mb-8 '>
									Quality freshly roasted coffee made just for you. Pour, brew
									and enjoy
								</p>
							</div>
							<div className='relative lg:w-4/12'>
								<div className=' h-[5rem] w-[20rem] absolute  lg:top-[14.5rem] right-16'>
									<img src={"/img/Waves.svg"} alt='' />
								</div>
								<div className='h-[10rem] w-[18rem]  sm:h-[16rem] sm:w-[26rem] lg:h-[17.8rem] lg:w-[28rem] absolute -top-8 sm:-top-14 md:-top-7 -right-14 lg:-right-36'>
									<img
										className='object-cover object-center h-full w-full'
										src={"/uploads/hero.png"}
										alt=''
									/>
								</div>
							</div>
						</div>
					</div>
				</section>
				<section className=''>
					{productDataIsLoading ? (
						<div className='flex h-[20rem] w-full items-center justify-center'>
							<b className='text-3xl'>Loading Your Product</b>
						</div>
					) : productData === "" || productData === undefined ? (
						state.user.role === "admin" ? (
							<>
								<div className='flex h-[20rem] w-full items-center justify-center'>
									<b className='text-3xl'>
										<Link
											to={"/product_add"}
											className='hover:text-coffee-300 transition-all'
										>
											Tambahkan Produk
										</Link>
									</b>
								</div>
							</>
						) : (
							<div className='flex h-[20rem] w-full items-center justify-center'>
								<b className='text-3xl'>Produk Kosong</b>
							</div>
						)
					) : (
						<div className='grid lg:grid-cols-4 gap-5 pt-36 px-4 lg:px-0 lg:pt-5 pb-5'>
							{productData?.map((data) => (
								<div
									key={data.id}
									className='bg-coffee-100 cursor-pointer hover:bg-coffee-200 transition-all'
									onClick={() => {
										navigate(`/detail/${data.id}`);
									}}
								>
									<img src={data.image} alt='' />
									<div className='p-2'>
										<h5 className='text-2xl font-bold tracking-tight text-coffee-400 '>
											{data.name}
										</h5>
										<p className='font-normal text-coffee-300'>
											{toCurrency(data.price)} <br />
											Stock : {data.stock}
										</p>
									</div>
								</div>
							))}
						</div>
					)}
				</section>
			</>
		</Layout>
	);
};

export { Home };
