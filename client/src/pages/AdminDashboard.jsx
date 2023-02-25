import React from "react";
import { Table } from "flowbite-react";

import Layout from "@/layouts/Default";

const AdminDashboard = () => {
	const title = "Dashboard";
	document.title = "WaysBeans | " + title;
	return (
		<Layout className='max-w-screen-xl mx-auto'>
			<>
				<section className='h-full mt-12 lg:mt-32 px-3 lg:px-6'>
					<h2 className='text-3xl mb-6 lg:mb-12 text-coffee-400 font-bold'>
						Income Transaction
					</h2>
					<div className='max-w-screen-lg mx-auto'>
						<Table striped={true}>
							<Table.Head>
								<Table.HeadCell className='bg-gray-200'>No</Table.HeadCell>
								<Table.HeadCell className='bg-gray-200'>
									Product name
								</Table.HeadCell>
								<Table.HeadCell className='bg-gray-200'>Color</Table.HeadCell>
								<Table.HeadCell className='bg-gray-200'>
									Category
								</Table.HeadCell>
								<Table.HeadCell className='bg-gray-200'>Price</Table.HeadCell>
								<Table.HeadCell className='bg-gray-200'>Status</Table.HeadCell>
							</Table.Head>
							<Table.Body className='divide-y'>
								<Table.Row className='bg-white dark:border-gray-700 dark:bg-gray-800'>
									<Table.Cell className='w-8'>1</Table.Cell>
									<Table.Cell className='whitespace-nowrap font-medium text-gray-900 w-full'>
										Apple MacBook Pro 17"
									</Table.Cell>
									<Table.Cell>Sliver</Table.Cell>
									<Table.Cell>Laptop</Table.Cell>
									<Table.Cell>$2999</Table.Cell>
									<Table.Cell>
										<a
											href='/tables'
											className='font-medium text-blue-600 hover:underline dark:text-blue-500'
										>
											Edit
										</a>
									</Table.Cell>
								</Table.Row>
								<Table.Row className='bg-white dark:border-gray-700 dark:bg-gray-800'>
									<Table.Cell className='w-8'>1</Table.Cell>
									<Table.Cell className='whitespace-nowrap font-medium text-gray-900 w-full'>
										Microsoft Surface Pro
									</Table.Cell>
									<Table.Cell>White</Table.Cell>
									<Table.Cell>Laptop PC</Table.Cell>
									<Table.Cell>$1999</Table.Cell>
									<Table.Cell>
										<a
											href='/tables'
											className='font-medium text-blue-600 hover:underline dark:text-blue-500'
										>
											Edit
										</a>
									</Table.Cell>
								</Table.Row>
								<Table.Row className='bg-white dark:border-gray-700 dark:bg-gray-800'>
									<Table.Cell className='w-8'>1</Table.Cell>
									<Table.Cell className='whitespace-nowrap font-medium text-gray-900 w-full'>
										Magic Mouse 2
									</Table.Cell>
									<Table.Cell>Black</Table.Cell>
									<Table.Cell>Accessories</Table.Cell>
									<Table.Cell>$99</Table.Cell>
									<Table.Cell>
										<a
											href='/tables'
											className='font-medium text-blue-600 hover:underline dark:text-blue-500'
										>
											Edit
										</a>
									</Table.Cell>
								</Table.Row>
								<Table.Row className='bg-white dark:border-gray-700 dark:bg-gray-800'>
									<Table.Cell className='w-8'>1</Table.Cell>
									<Table.Cell className='whitespace-nowrap font-medium text-gray-900 w-full'>
										Google Pixel Phone
									</Table.Cell>
									<Table.Cell>Gray</Table.Cell>
									<Table.Cell>Phone</Table.Cell>
									<Table.Cell>$799</Table.Cell>
									<Table.Cell>
										<a
											href='/tables'
											className='font-medium text-blue-600 hover:underline dark:text-blue-500'
										>
											Edit
										</a>
									</Table.Cell>
								</Table.Row>
								<Table.Row className='bg-white dark:border-gray-700 dark:bg-gray-800'>
									<Table.Cell className='w-8'>1</Table.Cell>
									<Table.Cell className='whitespace-nowrap font-medium text-gray-900 w-full'>
										Apple Watch 5
									</Table.Cell>
									<Table.Cell>Red</Table.Cell>
									<Table.Cell>Wearables</Table.Cell>
									<Table.Cell>$999</Table.Cell>
									<Table.Cell>
										<a
											href='/tables'
											className='font-medium text-blue-600 hover:underline dark:text-blue-500'
										>
											Edit
										</a>
									</Table.Cell>
								</Table.Row>
							</Table.Body>
						</Table>
					</div>
				</section>
			</>
		</Layout>
	);
};

export { AdminDashboard };
