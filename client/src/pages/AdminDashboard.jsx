import React from "react";
import { Table } from "flowbite-react";
import { useQuery } from "react-query";
import { API } from "@/lib/api";

import Layout from "@/layouts/Default";
import { toCurrency } from "@/lib/currency";

const AdminDashboard = () => {
	let { data: transactions, refetch } = useQuery(
		"transactionsCache",
		async () => {
			const response = await API.get("/transactions");
			return response.data.data;
		}
	);
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
									Customer Name
								</Table.HeadCell>
								<Table.HeadCell className='bg-gray-200'>
									Transaction ID
								</Table.HeadCell>
								<Table.HeadCell className='bg-gray-200'>Price</Table.HeadCell>
								<Table.HeadCell className='bg-gray-200'>Status</Table.HeadCell>
							</Table.Head>
							<Table.Body className='divide-y'>
								{transactions?.map((item, index) => (
									<Table.Row
										key={index}
										className='bg-white dark:border-gray-700 dark:bg-gray-800'
									>
										<Table.Cell className='w-8'>{index + 1}</Table.Cell>
										<Table.Cell className='whitespace-nowrap font-medium text-gray-900'>
											{item.user.name}
										</Table.Cell>
										<Table.Cell>{item.id}</Table.Cell>
										<Table.Cell>{toCurrency(item?.total)}</Table.Cell>
										<Table.Cell
											className={`font-medium capitalize ${
												item.status === "success"
													? "text-lime-400"
													: item.status === "failed"
													? "text-red-600"
													: item.status === "pending"
													? "text-amber-400"
													: item.status === "On the way"
													? "text-cyan-400"
													: ""
											}`}
										>
											{item.status}
										</Table.Cell>
									</Table.Row>
								))}
							</Table.Body>
						</Table>
					</div>
				</section>
			</>
		</Layout>
	);
};

export { AdminDashboard };
