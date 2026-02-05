import { type ReactNode } from 'react';

interface BadgeProps {
  children: ReactNode;
  /** Visual variant */
  variant?: 'default' | 'brand' | 'success' | 'warning' | 'danger';
  /** Badge size */
  size?: 'sm' | 'md';
  /** Optional dot indicator */
  dot?: boolean;
  /** Dot color (when dot=true) */
  dotColor?: 'brand' | 'success' | 'danger';
  /** Additional classes */
  className?: string;
}

const variantClasses = {
  default: 'bg-slate-100 text-slate-600 dark:bg-zinc-800 dark:text-zinc-400',
  brand: 'bg-blue-100 text-[#1e3a8a] dark:bg-blue-900/50 dark:text-blue-300',
  success: 'bg-emerald-100 text-emerald-700 dark:bg-emerald-900/50 dark:text-emerald-400',
  warning: 'bg-amber-100 text-amber-700 dark:bg-amber-900/50 dark:text-amber-400',
  danger: 'bg-red-100 text-red-700 dark:bg-red-900/50 dark:text-red-400',
} as const;

const sizeClasses = {
  sm: 'px-2 py-0.5 text-[10px]',
  md: 'px-2.5 py-1 text-xs',
} as const;

const dotColorClasses = {
  brand: 'bg-[#1e3a8a]',
  success: 'bg-emerald-500',
  danger: 'bg-[#dc2626]',
} as const;

export function Badge({
  children,
  variant = 'default',
  size = 'sm',
  dot = false,
  dotColor = 'brand',
  className = '',
}: BadgeProps) {
  return (
    <span
      className={`inline-flex items-center gap-1.5 rounded-full font-semibold ${variantClasses[variant]} ${sizeClasses[size]} ${className}`}
    >
      {dot && <span className={`h-1.5 w-1.5 rounded-full ${dotColorClasses[dotColor]}`} />}
      {children}
    </span>
  );
}
