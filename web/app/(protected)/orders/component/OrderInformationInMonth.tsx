'use client'

import { Bar, BarChart, CartesianGrid, XAxis } from "recharts"
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import {
    ChartConfig,
    ChartContainer,
    ChartTooltip,
    ChartTooltipContent,
} from "@/components/ui/chart"

const chartData = [
    { month: "T1", orders: 186, revenue: 8000000 },
    { month: "T2", orders: 305, revenue: 12000000 },
    { month: "T3", orders: 237, revenue: 9500000 },
    { month: "T4", orders: 73, revenue: 3200000 },
    { month: "T5", orders: 209, revenue: 11000000 },
    { month: "T6", orders: 214, revenue: 10500000 },
]

const chartConfig = {
    orders: {
        label: "Đơn hàng",
        color: "hsl(var(--chart-1))",
    },
    revenue: {
        label: "Doanh thu",
        color: "hsl(var(--chart-2))",
    },
} satisfies ChartConfig

export default function OrderInformationInMonth() {

    return (
            <div>
                <h2 className='text-lg font-semibold'>Tổng quan</h2>

                <div className='flex gap-8 mt-4 *:basis-1/3'>
                    <Card>
                        <CardContent className='flex flex-col items-center gap-4 justify-center h-full'>
                            <span className='text-4xl font-semibold'>40</span>
                            <span>đơn hàng trong tháng</span>
                        </CardContent>

                    </Card>

                    <Card>
                        <CardContent className='flex flex-col items-center gap-4 justify-center h-full'>
                            <span className='text-2xl font-semibold'>12,000,000đ</span>
                            <span>doanh thu trong tháng</span>
                        </CardContent>
                    </Card>

                    <div>
                        <Card>
                            <CardHeader>
                                <CardTitle>Biểu đồ đơn hàng theo tháng</CardTitle>
                                <CardDescription>Thống kê 6 tháng gần nhất</CardDescription>
                            </CardHeader>
                            <CardContent>
                                <ChartContainer config={chartConfig}>
                                    <BarChart accessibilityLayer data={chartData}>
                                        <CartesianGrid vertical={false} />
                                        <XAxis
                                            dataKey="month"
                                            tickLine={false}
                                            tickMargin={10}
                                            axisLine={false}
                                        />
                                        <ChartTooltip
                                            cursor={false}
                                            content={<ChartTooltipContent indicator="dashed" />}
                                        />
                                        <Bar dataKey="orders" fill="var(--color-orders)" radius={4} />
                                        <Bar dataKey="revenue" fill="var(--color-revenue)" radius={4} />
                                    </BarChart>
                                </ChartContainer>
                            </CardContent>
                        </Card>
                    </div>

                </div>


            </div>
    )
}