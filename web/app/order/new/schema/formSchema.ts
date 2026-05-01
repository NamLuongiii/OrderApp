import z from "zod";

const checkoutSchema = z.object({
    customerName: z.string().min(2, "Tên phải ít nhất 2 ký tự"),
    email: z.string().email("Email không hợp lệ"),
    phone: z.string().min(10, "Số điện thoại không hợp lệ"),
    address: z.string().min(5, "Địa chỉ quá ngắn"),
    note: z.string().optional(),
});

type CheckoutFormData = z.infer<typeof checkoutSchema>;

export { checkoutSchema, type CheckoutFormData };