import { Link } from '@inertiajs/react';
import { ExternalLink, type LucideIcon } from 'lucide-react';

interface NavItemProps {
  /** Display name */
  name: string;
  /** Navigation URL */
  href: string;
  /** Icon component */
  icon: LucideIcon;
  /** Is this item active? */
  active?: boolean;
  /** Is this an external link? */
  external?: boolean;
  /** Click handler */
  onClick?: () => void;
  /** Additional classes */
  className?: string;
}

export function NavItem({
  name,
  href,
  icon: Icon,
  active = false,
  external = false,
  onClick,
  className = '',
}: NavItemProps) {
  const baseClasses = 'flex items-center gap-3 rounded-xl px-3 py-2.5 text-sm font-medium transition-all duration-150 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-[#1e3a8a]/30 focus-visible:ring-offset-2 dark:focus-visible:ring-offset-zinc-900';
  const activeClasses = active
    ? 'bg-[#1e3a8a] text-white shadow-sm'
    : 'text-slate-600 hover:bg-slate-100 dark:text-zinc-400 dark:hover:bg-zinc-800';

  if (external) {
    return (
      <a
        href={href}
        target="_blank"
        rel="noopener noreferrer"
        className={`${baseClasses} ${activeClasses} ${className}`}
        onClick={onClick}
      >
        <Icon className="h-5 w-5" />
        {name}
        <ExternalLink className="ml-auto h-4 w-4 opacity-50" />
      </a>
    );
  }

  return (
    <Link
      href={href}
      className={`${baseClasses} ${activeClasses} ${className}`}
      onClick={onClick}
    >
      <Icon className="h-5 w-5" />
      {name}
    </Link>
  );
}
