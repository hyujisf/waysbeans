import React, { useEffect, useState } from "react";
import { Table } from "flowbite-react";
import { useNavigate } from "react-router-dom";
import { useMutation, useQuery } from "react-query";
import { Button, Spinner } from "flowbite-react";
import Swal from "sweetalert2";

import { API } from "@/lib/api";
import { toCurrency } from "@/lib/currency";

import Layout from "@/layouts/Default";

const AdminListProduct = () => {
	let { data: products, refetch } = useQuery("productCache", async () => {
		const response = await API.get("/products");
		return response.data.data;
	});

	//Update Product
	const handleUpdate = (id) => {
		navigate("/update-product/" + id);
	};

	const handleDeleteProduct = useMutation(async (id) => {
		try {
			await Swal.fire({
				icon: "warning",
				title: "Do you want to remove this product? ?",
				confirmButtonText: "Yes",
				cancelButtonText: "No",
				showCancelButton: true,
				confirmButtonColor: "#d33",
				cancelButtonColor: "rgb(97 61 43)",
			}).then((result) => {
				/* Read more about isConfirmed, isDenied below */
				if (result.isConfirmed) {
					const response = API.delete(`/product/${id}`);
					refetch();
					Swal.fire("Berhasil dihapus!", "", "success");
				}
			});
		} catch (e) {
			console.log(e);
		}
	});
	const title = "List Product";
	document.title = "WaysBeans | " + title;
	return (
		<Layout className='max-w-screen-xl mx-auto'>
			<>
				<section className='h-full mt-12 lg:mt-32 px-3 lg:px-6'>
					<h2 className='text-3xl mb-6 lg:mb-12 text-coffee-400 font-bold'>
						Income Transaction
					</h2>
					<Table striped={true}>
						<Table.Head>
							<Table.HeadCell className='bg-gray-200'>No</Table.HeadCell>
							<Table.HeadCell className='bg-gray-200'>Image</Table.HeadCell>
							<Table.HeadCell className='bg-gray-200'>Name</Table.HeadCell>
							<Table.HeadCell className='bg-gray-200'>Stock</Table.HeadCell>
							<Table.HeadCell className='bg-gray-200'>Price</Table.HeadCell>
							<Table.HeadCell className='bg-gray-200'>
								Description
							</Table.HeadCell>
							<Table.HeadCell className='bg-gray-200'>Status</Table.HeadCell>
						</Table.Head>
						<Table.Body className='divide-y'>
							{products != null ? (
								products?.map((product, k) => (
									<Table.Row
										key={k}
										className='bg-white dark:border-gray-700 dark:bg-gray-800'
									>
										<Table.Cell className='w-8'> {k + 1}</Table.Cell>
										<Table.Cell>
											<img src={product?.image} className='w-28' alt='' />
										</Table.Cell>
										<Table.Cell className='whitespace-nowrap font-medium text-gray-900'>
											{product.name}
										</Table.Cell>
										<Table.Cell>{product.stock}</Table.Cell>
										<Table.Cell>{toCurrency(product.price)}</Table.Cell>
										<Table.Cell className='max-w-[25rem] text-justify'>
											{product.description}
										</Table.Cell>
										<Table.Cell>
											<div className='flex gap-4'>
												{handleDeleteProduct.isLoading ? (
													<Button
														color='failure'
														onClick={() => {
															handleDeleteProduct.mutate(product.id);
														}}
														disabled
													>
														<Spinner aria-label='Delete Spinner' />
													</Button>
												) : (
													<Button
														color='failure'
														onClick={() => {
															handleDeleteProduct.mutate(product.id);
														}}
													>
														Delete
													</Button>
												)}
												<Button
													color='success'
													onClick={() => {
														handleUpdate.mutate(product.id);
													}}
													type='button'
													className='buttonList'
												>
													Update
												</Button>
											</div>
										</Table.Cell>
									</Table.Row>
								))
							) : (
								<Table.Row>
									<Table.Cell
										colSpan={7}
										className='text-center text-coffee-300 font-medium text-2xl py-9'
									>
										Produk kosong
									</Table.Cell>
								</Table.Row>
							)}
						</Table.Body>
					</Table>
				</section>
			</>
		</Layout>
	);
};

export { AdminListProduct };
