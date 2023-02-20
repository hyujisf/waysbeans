import React from "react";
import { Link } from "react-router-dom";
import { useParams } from "react-router-dom";
import { Button } from "flowbite-react";

import Layout from "@/layouts/Default";

const Detail = () => {
	const { id } = useParams();
	const coffeeData = [
		{
			img: "/uploads/product1.png",
			name: "RWANDA Beans",
			price: "Rp. 299.000",
			stock: "200",
		},
		{
			img: "/uploads/product2.png",
			name: "RWANDA Beans",
			price: "Rp. 299.000",
			stock: "200",
		},
		{
			img: "/uploads/product3.png",
			name: "RWANDA Beans",
			price: "Rp. 299.000",
			stock: "200",
		},
		{
			img: "/uploads/product4.png",
			name: "RWANDA Beans",
			price: "Rp. 299.000",
			stock: "200",
		},
	];
	return (
		<Layout className='max-w-screen-lg mx-auto'>
			<>
				<section className='h-full mt-6 lg:mt-32 lg:px-6'>
					<div className='flex flex-col lg:flex-row my-auto'>
						<div className='lg:w-[46rem]'>
							<img
								src={coffeeData[id].img}
								className={"object-cover object-center h-full w-full"}
								alt=''
							/>
						</div>
						<div className='w-full p-12'>
							<h1 className='text-5xl font-bold text-coffee-400'>
								{coffeeData[id].name}
							</h1>
							<p className='mt-3 text-coffee-300'>
								Stock {coffeeData[id].stock}
							</p>
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
								{coffeeData[id].price}
							</p>

							<Button
								type='button'
								size={"lg"}
								className='bg-coffee-400 hover:bg-transparent border-2 !border-coffee-400 text-white hover:text-coffee-400 px-4 py-0 font-semibold transition-all !w-full md:w-auto'
							>
								Add Cart
							</Button>
						</div>
					</div>
				</section>
			</>
		</Layout>
	);
};

export { Detail };
