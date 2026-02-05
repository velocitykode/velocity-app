import { type LucideIcon } from 'lucide-react';
import { type ComponentProps, forwardRef } from 'react';

type ButtonOrLinkProps =
  | (ComponentProps<'button'> & { as?: 'button'; href?: never })
  | (ComponentProps<'a'> & { as: 'a'; href: string });

interface IconButtonBaseProps {
  /** Lucide icon component */
  icon: LucideIcon;
  /** Button size */
  size?: 'sm' | 'md' | 'lg';
  /** Visual variant */
  variant?: 'default' | 'danger' | 'ghost';
  /** Accessible label */
  'aria-label': string;
}

type IconButtonProps = IconButtonBaseProps & ButtonOrLinkProps;

const sizeClasses = {
  sm: 'h-7 w-7',
  md: 'h-8 w-8',
  lg: 'h-10 w-10',
} as const;

const iconSizeClasses = {
  sm: 'h-3.5 w-3.5',
  md: 'h-4 w-4',
  lg: 'h-5 w-5',
} as const;

const variantClasses = {
  default: 'bg-slate-100 text-slate-500 hover:bg-slate-200 hover:text-slate-700 dark:bg-zinc-800 dark:text-zinc-400 dark:hover:bg-zinc-700 dark:hover:text-zinc-200 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-[#1e3a8a]/30 focus-visible:ring-offset-2 dark:focus-visible:ring-offset-zinc-900',
  danger: 'bg-slate-100 text-slate-500 hover:bg-red-100 hover:text-red-600 dark:bg-zinc-800 dark:text-zinc-400 dark:hover:bg-red-950/50 dark:hover:text-red-400 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-red-500/30 focus-visible:ring-offset-2 dark:focus-visible:ring-offset-zinc-900',
  ghost: 'text-slate-500 hover:bg-slate-100 hover:text-slate-700 dark:text-zinc-400 dark:hover:bg-zinc-800 dark:hover:text-zinc-200 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-[#1e3a8a]/30 focus-visible:ring-offset-2 dark:focus-visible:ring-offset-zinc-900',
} as const;

export const IconButton = forwardRef<HTMLButtonElement | HTMLAnchorElement, IconButtonProps>(
  ({ icon: Icon, size = 'md', variant = 'default', className = '', as, ...props }, ref) => {
    const classes = `flex items-center justify-center rounded-xl transition-all duration-150 active:scale-95 ${sizeClasses[size]} ${variantClasses[variant]} ${className}`;

    if (as === 'a') {
      const { href, ...anchorProps } = props as ComponentProps<'a'> & { href: string };
      return (
        <a ref={ref as React.Ref<HTMLAnchorElement>} href={href} className={classes} {...anchorProps}>
          <Icon className={iconSizeClasses[size]} />
        </a>
      );
    }

    return (
      <button ref={ref as React.Ref<HTMLButtonElement>} className={classes} {...(props as ComponentProps<'button'>)}>
        <Icon className={iconSizeClasses[size]} />
      </button>
    );
  }
);

IconButton.displayName = 'IconButton';
