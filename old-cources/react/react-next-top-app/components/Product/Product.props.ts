import { DetailedHTMLProps, HTMLAttributes } from "react";
import { IProductModel } from "../../interfaces/product.interface";

export interface ProductProps extends DetailedHTMLProps<HTMLAttributes<HTMLDivElement>, HTMLDivElement> {
	product: IProductModel
}