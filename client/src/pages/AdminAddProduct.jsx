import React, { useState } from "react";
import { Link } from "react-router-dom";
import { useParams } from "react-router-dom";
import { useMutation } from "react-query";
import { Button, Label, TextInput } from "flowbite-react";
import { API } from "@/lib/api";
import { useQuery } from "react-query";

import Toast from "@/lib/alert";
import Layout from "@/layouts/Default";

const AdminAddProduct = () => {
	const [preview, setPreview] = useState(null);
	const [inputProduct, setInputProduct] = useState({
		name: "",
		price: "",
		stock: "",
		description: "",
		image: "",
	});

	// Handle change data on form
	const handleChange = (e) => {
		setInputProduct({
			...inputProduct,
			[e.target.name]:
				e.target.type === "file" ? e.target.files : e.target.value,
		});

		// Create image url for preview
		if (e.target.type === "file") {
			let url = URL.createObjectURL(e.target.files[0]);
			setPreview(url);
		}
	};
	const handleSubmit = useMutation(async (e) => {
		try {
			e.preventDefault();

			// Configuration
			const config = {
				headers: {
					"Content-type": "multipart/form-data",
				},
			};
			const data = new FormData();
			data.append("name", inputProduct.name);
			data.append("price", inputProduct.price);
			data.append("stock", inputProduct.stock);
			data.append("description", inputProduct.description);
			data.append("image", inputProduct.image[0]);

			await API.post("/product", data, config);

			e.target.reset();
			setPreview(null);
			Toast.fire({
				icon: "success",
				title: "Product success to add",
			});
			// redirect("/");
		} catch (err) {
			// console.log(form.amenities);

			Toast.fire({
				icon: "error",
				title: "Product fail to add",
			});
		}
	});

	const title = "Add Product";
	document.title = "Housy | " + title;

	return (
		<Layout className='max-w-screen-lg mx-auto'>
			<>
				<section className='h-full mt-6 lg:mt-32 lg:px-6'>
					<div className='flex flex-col lg:flex-row my-auto'>
						<div className='w-full px-12'>
							<h1 className='text-3xl font-bold text-coffee-400'>
								Add Product{" "}
							</h1>
							<form
								onSubmit={(e) => handleSubmit.mutate(e)}
								className='flex flex-col gap-4 mt-10'
							>
								<div>
									<input
										type='text'
										name='name'
										placeholder='Name'
										onChange={handleChange}
										required
										className='w-full rounded-lg bg-coffee-400/25 placeholder:text-coffee-300 text-coffee-400 border-2 border-coffee-400'
									/>
								</div>
								<div>
									<input
										type='text'
										name='stock'
										placeholder='Stock'
										onChange={handleChange}
										required
										className='w-full rounded-lg bg-coffee-400/25 placeholder:text-coffee-300 text-coffee-400 border-2 border-coffee-400'
									/>
								</div>
								<div>
									<input
										type='text'
										name='price'
										placeholder='Price'
										onChange={handleChange}
										required
										className='w-full rounded-lg bg-coffee-400/25 placeholder:text-coffee-300 text-coffee-400 border-2 border-coffee-400'
									/>
								</div>
								<div>
									<textarea
										name='description'
										placeholder='Description Product'
										onChange={handleChange}
										required
										rows={5}
										className='w-full rounded-lg bg-coffee-400/25 placeholder:text-coffee-300 text-coffee-400 border-2 border-coffee-400'
									/>
								</div>
								<div>
									<input
										type='file'
										name='image'
										onChange={handleChange}
										required
										className='w-full rounded-lg bg-coffee-400/25 placeholder:text-coffee-300 text-coffee-400 border-2 border-coffee-400 cursor-text file:hidden px-3 py-2'
									/>
								</div>
								<Button
									type='submit'
									size={"lg"}
									className='bg-coffee-400 hover:bg-transparent border-2 !border-coffee-400 text-white hover:text-coffee-400 px-4 py-0 font-semibold transition-all !w-full'
								>
									Add Cart
								</Button>
							</form>
						</div>
						<div className='lg:w-[46rem]'>
							{preview && (
								<img
									src={preview}
									className={"object-cover object-center h-full w-full"}
									alt={preview}
								/>
							)}
						</div>
					</div>
				</section>
			</>
		</Layout>
	);
};

export { AdminAddProduct };
